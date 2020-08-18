package program

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// BasicProgram blabla
type BasicProgram struct {
	GlProgram uint32

	uModelToProjection int32
}

// NewBasicProgram blabla
func NewBasicProgram() (*BasicProgram, error) {
	var result BasicProgram

	glProgram, err := buildProgram(vertexShaderSource, fragmentShaderSource)
	if err != nil {
		return nil, err
	}

	result.GlProgram = glProgram
	result.uModelToProjection = gl.GetUniformLocation(glProgram, gl.Str("u_modelToProjection\x00"))

	return &result, nil
}

// Use blabla
func (p *BasicProgram) Use() {
	gl.UseProgram(p.GlProgram)
}

// SetUniformModelToProjection blabla
func (p *BasicProgram) SetUniformModelToProjection(modelToProjection *mgl32.Mat4) {
	gl.UniformMatrix4fv(p.uModelToProjection, 1, false, &modelToProjection[0])
}

var vertexShaderSource = `
#version 410

uniform mat4 u_modelToProjection;

in vec4 in_position;

void main() {
  gl_Position = u_modelToProjection * in_position;
}
` + "\x00"

var fragmentShaderSource = `
#version 410

out vec4 out_color;

void main() {
  out_color = vec4(1.0, 0.0, 0.0, 1.0);
}
` + "\x00"
