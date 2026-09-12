package mbtiles

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
)

func validateMergeMetadata(metadata map[string]string) error {
	if strings.TrimSpace(metadata["name"]) == "" || strings.TrimSpace(metadata["format"]) == "" {
		return fmt.Errorf("metadata requires name and format")
	}
	if scheme := metadata["scheme"]; scheme != "" && scheme != "tms" {
		return fmt.Errorf("unsupported MBTiles scheme: %q", scheme)
	}
	if bounds := metadata["bounds"]; bounds != "" {
		if _, err := parseMergeBounds(bounds); err != nil {
			return err
		}
	}
	if metadata["format"] == "pbf" {
		object, err := parseMetadataJSON(metadata["json"])
		if err != nil {
			return err
		}
		if _, ok := object["vector_layers"]; !ok {
			return fmt.Errorf("pbf metadata requires json.vector_layers")
		}
	}
	return nil
}

func parseMergeBounds(value string) ([]float64, error) {
	bounds := convertPoints(strings.Split(value, ","))
	if len(bounds) != 4 || bounds[0] > bounds[2] || bounds[1] > bounds[3] || bounds[0] < -180 || bounds[2] > 180 || bounds[1] < -90 || bounds[3] > 90 {
		return nil, fmt.Errorf("invalid metadata bounds: %q", value)
	}
	return bounds, nil
}

func combineMergeMetadata(dest, src map[string]string) error {
	if dest["format"] != "" && dest["format"] != src["format"] {
		return fmt.Errorf("incompatible tile formats: %q and %q", dest["format"], src["format"])
	}
	for name, value := range src {
		switch name {
		case "bounds":
			incoming, err := parseMergeBounds(value)
			if err != nil {
				return err
			}
			if previous := dest[name]; previous != "" {
				existing, err := parseMergeBounds(previous)
				if err != nil {
					return err
				}
				incoming[0] = math.Min(incoming[0], existing[0])
				incoming[1] = math.Min(incoming[1], existing[1])
				incoming[2] = math.Max(incoming[2], existing[2])
				incoming[3] = math.Max(incoming[3], existing[3])
				value = fmt.Sprintf("%g,%g,%g,%g", incoming[0], incoming[1], incoming[2], incoming[3])
			}
		case "json":
			if previous := dest[name]; previous != "" {
				merged, err := mergeMetadataJSON(previous, value)
				if err != nil {
					return err
				}
				value = merged
			}
		}
		dest[name] = value
	}
	return nil
}

func parseMetadataJSON(value string) (map[string]json.RawMessage, error) {
	var object map[string]json.RawMessage
	if err := json.Unmarshal([]byte(value), &object); err != nil {
		return nil, fmt.Errorf("invalid metadata json: %w", err)
	}
	if object == nil {
		return nil, fmt.Errorf("metadata json must be an object")
	}
	if raw, ok := object["vector_layers"]; ok {
		if _, err := parseVectorLayers(raw); err != nil {
			return nil, err
		}
	}
	return object, nil
}

func parseVectorLayers(raw json.RawMessage) ([]map[string]json.RawMessage, error) {
	var layers []map[string]json.RawMessage
	if err := json.Unmarshal(raw, &layers); err != nil {
		return nil, fmt.Errorf("invalid vector_layers: %w", err)
	}
	if layers == nil {
		return nil, fmt.Errorf("vector_layers must be an array")
	}
	for _, layer := range layers {
		var id string
		if err := json.Unmarshal(layer["id"], &id); err != nil || id == "" {
			return nil, fmt.Errorf("vector layer requires a string id")
		}
		var fields map[string]string
		if err := json.Unmarshal(layer["fields"], &fields); err != nil || fields == nil {
			return nil, fmt.Errorf("vector layer %q requires fields", id)
		}
	}
	return layers, nil
}

// Preserve extension keys, combine vector layers by ID and union their fields.
// Conflicting field types are rejected instead of emitting misleading metadata.
func mergeMetadataJSON(left, right string) (string, error) {
	dest, err := parseMetadataJSON(left)
	if err != nil {
		return "", err
	}
	src, err := parseMetadataJSON(right)
	if err != nil {
		return "", err
	}
	for key, value := range src {
		if key != "vector_layers" || dest[key] == nil {
			dest[key] = value
			continue
		}
		existing, err := parseVectorLayers(dest[key])
		if err != nil {
			return "", err
		}
		incoming, err := parseVectorLayers(value)
		if err != nil {
			return "", err
		}
		indices := make(map[string]int)
		for i, layer := range existing {
			var id string
			_ = json.Unmarshal(layer["id"], &id) // Validated by parseVectorLayers.
			indices[id] = i
		}
		for _, layer := range incoming {
			var id string
			_ = json.Unmarshal(layer["id"], &id) // Validated by parseVectorLayers.
			index, ok := indices[id]
			if !ok {
				indices[id] = len(existing)
				existing = append(existing, layer)
				continue
			}
			target := existing[index]
			for name, raw := range layer {
				switch name {
				case "fields":
					var a, b map[string]string
					if err := json.Unmarshal(target[name], &a); err != nil {
						return "", err
					}
					if err := json.Unmarshal(raw, &b); err != nil {
						return "", err
					}
					for field, fieldType := range b {
						if previous, ok := a[field]; ok && previous != fieldType {
							return "", fmt.Errorf("incompatible vector field types in layer %s for %q", id, field)
						}
						a[field] = fieldType
					}
					raw, err = json.Marshal(a)
					if err != nil {
						return "", err
					}
				case "minzoom", "maxzoom":
					if target[name] != nil {
						var a, b float64
						if err := json.Unmarshal(target[name], &a); err != nil {
							return "", err
						}
						if err := json.Unmarshal(raw, &b); err != nil {
							return "", err
						}
						if name == "minzoom" {
							b = math.Min(a, b)
						} else {
							b = math.Max(a, b)
						}
						raw, err = json.Marshal(b)
						if err != nil {
							return "", err
						}
					}
				}
				target[name] = raw
			}
		}
		dest[key], err = json.Marshal(existing)
		if err != nil {
			return "", err
		}
	}
	data, err := json.Marshal(dest)
	return string(data), err
}
