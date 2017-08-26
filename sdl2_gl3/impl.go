package imgui_sdl

import (
	"github.com/armored-dragon/go-imgui"
	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	vertexShader = `#version 330
uniform mat4 ProjMtx;
in vec2 Position;
in vec2 UV;
in vec4 Color;
out vec2 Frag_UV;
out vec4 Frag_Color;
void main() {
    Frag_UV = UV;
    Frag_Color = Color;
    gl_Position = ProjMtx * vec4(Position.xy, 0, 1);
}`
	fragmentShader = `#version 330
uniform sampler2D Texture;
in vec2 Frag_UV;
in vec4 Frag_Color;
out vec4 Out_Color;
void main() {
    Out_Color = Frag_Color * texture( Texture, Frag_UV.st);
}`
)

var (
	time                   float64 = 0.0
	mousePressed                   = []bool{false, false, false}
	mouseWheel             float32 = 0.0
	fontTexture            uint32  = 0 // ????????? GLuint
	shaderHandle           uint32  = 0
	vertHandle             int32   = 0
	fragHandle             int32   = 0
	attribLocationTex      int32   = 0
	attribLocationProjMtx  int32   = 0
	attribLocationPosition int32   = 0
	attribLocationUV       int32   = 0
	attribLocationColor    int32   = 0
	vboHandle              uint32  = 0
	vaoHandle              uint32  = 0
	elementsHandle         uint32  = 0
)

func getIntegerv(pname uint32) int32 {
	var value int32
	gl.GetIntegerv(pname, &value)
	return value
}

func getIntegerv4(target uint32) [4]int32 {
	var values [4]int32
	gl.GetIntegeri_v(target, 0, &values[0])
	gl.GetIntegeri_v(target, 1, &values[1])
	gl.GetIntegeri_v(target, 2, &values[2])
	gl.GetIntegeri_v(target, 3, &values[3])
	return values
}

func RenderDrawLists(drawData *imgui.DrawData) {
	io := imgui.GetIO()
	fbWidth := int32(io.GetDisplaySize().GetX() * io.GetDisplayFramebufferScale().GetX())
	fbHeight := int32(io.GetDisplaySize().GetY() * io.GetDisplayFramebufferScale().GetY())
	if fbWidth == 0 || fbHeight == 0 {
		return
	}
	drawData.ScaleClipRects(io.GetDisplayFramebufferScale())

	oldActiveTexture := getIntegerv(gl.ACTIVE_TEXTURE)
	gl.ActiveTexture(gl.TEXTURE0)
	oldProgram := getIntegerv(gl.CURRENT_PROGRAM)
	oldTexture := getIntegerv(gl.TEXTURE_BINDING_2D)
	oldArrayBuffer := getIntegerv(gl.ARRAY_BUFFER_BINDING)
	oldElementArrayBuffer := getIntegerv(gl.ELEMENT_ARRAY_BUFFER_BINDING)
	oldVertexArray := getIntegerv(gl.VERTEX_ARRAY_BINDING)
	oldViewport := getIntegerv4(gl.VIEWPORT)
	oldScissorBox := getIntegerv4(gl.SCISSOR_BOX)
	oldBlendSrcRgb := getIntegerv(gl.BLEND_SRC_RGB)
	oldBlendDstRgb := getIntegerv(gl.BLEND_DST_RGB)
	oldBlendSrcAlpha := getIntegerv(gl.BLEND_SRC_ALPHA)
	oldBlendDstAlpha := getIntegerv(gl.BLEND_DST_ALPHA)
	oldBlendEquationRgb := getIntegerv(gl.BLEND_EQUATION_RGB)
	oldBlendEquationAlpha := getIntegerv(gl.BLEND_EQUATION_ALPHA)
	oldEnableBlend := gl.IsEnabled(gl.BLEND)
	oldEnableCullFace := gl.IsEnabled(gl.CULL_FACE)
	oldEnableDepthTest := gl.IsEnabled(gl.DEPTH_TEST)
	oldEnableScissorTest := gl.IsEnabled(gl.SCISSOR_TEST)

	gl.Enable(gl.BLEND)
	gl.BlendEquation(gl.FUNC_ADD)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Disable(gl.CULL_FACE)
	gl.Disable(gl.DEPTH_TEST)
	gl.Enable(gl.SCISSOR_TEST)

	gl.Viewport(0, 0, fbWidth, fbHeight)
	orthoProjection := [4][4]float32{
		[4]float32{2.0 / io.GetDisplaySize().GetX(), 0.0, 0.0, 0.0},
		[4]float32{0.0, 2.0 / -io.GetDisplaySize().GetY(), 0.0, 0.0},
		[4]float32{0.0, 0.0, -1.0, 0.0},
		[4]float32{-1.0, 1.0, 0.0, 1.0},
	}
	gl.UseProgram(shaderHandle)
	gl.Uniform1i(attribLocationTex, 0)
	gl.UniformMatrix4fv(attribLocationProjMtx, 1, false, &orthoProjection[0][0])
	gl.BindVertexArray(vaoHandle)

	for _, cmdList := range drawData.GetCmdLists() {
		// TODO: implement the rest
	}
}
