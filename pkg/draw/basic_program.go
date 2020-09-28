package draw

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type BasicProgram struct {
	GlProgram uint32

	uModelToProjection           int32
	uModelToViewInverseTranspose int32
}

func (p *BasicProgram) Init() error {
	glProgram, err := buildProgram(vertexShaderSource, fragmentShaderSource, func(glProgram uint32) {
		gl.BindAttribLocation(glProgram, PositionAttributeLocation, gl.Str("in_position\x00"))
		gl.BindAttribLocation(glProgram, NormalAttributeLocation, gl.Str("in_normal\x00"))
	})
	if err != nil {
		return err
	}

	p.GlProgram = glProgram
	p.uModelToProjection = gl.GetUniformLocation(glProgram, gl.Str("u_modelToProjection\x00"))
	p.uModelToViewInverseTranspose = gl.GetUniformLocation(glProgram, gl.Str("u_modelToViewInverseTranspose\x00"))

	return nil
}

func (p *BasicProgram) Use() {
	gl.UseProgram(p.GlProgram)
}

func (p *BasicProgram) SetUniformModelToProjection(modelToProjection *mgl32.Mat4) {
	gl.UniformMatrix4fv(p.uModelToProjection, 1, false, &modelToProjection[0])
}

func (p *BasicProgram) SetUniformModelToViewInverseTranspose(modelToViewInverseTranspose *mgl32.Mat4) {
	gl.UniformMatrix4fv(p.uModelToViewInverseTranspose, 1, false, &modelToViewInverseTranspose[0])
}

var vertexShaderSource = `
#version 410

uniform mat4 u_modelToProjection;
uniform mat4 u_modelToViewInverseTranspose;

in vec4 in_position;
in vec3 in_normal;

out vec3 v_normal;

void main() {
	gl_Position = u_modelToProjection * in_position;
	v_normal = mat3(u_modelToViewInverseTranspose) * in_normal;
}
` + "\x00"

var fragmentShaderSource = `
#version 410

in vec3 v_normal;

out vec4 out_color;

void main() {
	float lightContribution = dot(normalize(v_normal), -1.0 * normalize(vec3(-0.5, -0.5, -0.5)));

  out_color = vec4(1.0, 0.0, 0.0, 1.0);
  out_color.rgb *= lightContribution;
}
` + "\x00"
