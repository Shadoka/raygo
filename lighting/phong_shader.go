package lighting

import (
	gomath "math"
	g "raygo/geometry"
	"raygo/math"
)

func PhongLighting(m g.Material, obj g.Shape, light Light, position math.Point, eyev math.Vector, normalv math.Vector, inShadow bool) math.Color {
	color := m.Color
	if m.Texture.Exists() {
		u, v := obj.GetUvCoordinate(normalv)
		color = m.Texture.ColorAt(u, v)
	} else if m.Pattern != nil {
		color = m.Pattern.ColorAtObject(position, obj)
	}

	// combine the surface color with the light's color/intensity
	effectiveColor := color.Blend(light.Intensity)

	// find the direction to the light source
	lightv := light.Position.Subtract(position).Normalize()

	// compute the ambient contribution
	ambient := effectiveColor.Mul(m.Ambient)

	if inShadow {
		return ambient
	}

	// lightDotNormal represents the cosine of the angle between the
	// light vector and the normal vector. A negative number means the
	// light is on the other side of the surface.
	lightDotNormal := lightv.Dot(normalv)
	diffuse := math.CreateColor(0.0, 0.0, 0.0)
	specular := math.CreateColor(0.0, 0.0, 0.0)
	if lightDotNormal >= 0 {
		// compute the diffuse contribution
		diffuse = effectiveColor.Mul(m.Diffuse).Mul(lightDotNormal)

		// reflectDotEye represents the cosine of the angle between the
		// reflection vector and the eye vector. A negative number means the
		// light reflects away from the eye
		reflectv := lightv.Negate().Reflect(normalv)
		reflectDotEye := reflectv.Dot(eyev)

		if reflectDotEye > 0 {
			// compute the specular contribution
			factor := gomath.Pow(reflectDotEye, m.Shininess)
			specular = light.Intensity.Mul(m.Specular).Mul(factor)
		}
	}

	return ambient.Add(diffuse).Add(specular)
}
