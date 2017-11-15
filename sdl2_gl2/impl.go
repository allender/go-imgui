package imgui_sdl_gl2

import (
	"github.com/allender/go-imgui"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	time                   float64 = 0.0
	mousePressed                   = []bool{false, false, false}
	mouseWheel             float32 = 0.0
	fontTexture            uint32  = 0 // ????????? GLuint
)

func RenderDrawLists(drawData *imgui.DrawData) {
	// Avoid rendering when minimized, scale coordinates for retina displays (screen coordinates != framebuffer coordinates)
	io := imgui.GetIO();
	fbWidth := int32(io.GetDisplaySize().GetX() * io.GetDisplayFramebufferScale().GetX())
	fbHeight := int32(io.GetDisplaySize().GetY() * io.GetDisplayFramebufferScale().GetY())
	if fbWidth == 0 || fbHeight == 0 {
		return
	}
	(*drawData).ScaleClipRects(io.GetDisplayFramebufferScale())

	// We are using the OpenGL fixed pipeline to make the example code simpler to read!
	// Setup render state: alpha-blending enabled, no face culling, no depth testing, scissor enabled, vertex/texcoord/color pointers.
	var last_texture int32
	var last_viewport [4]int32

	gl.GetIntegerv(gl.TEXTURE_BINDING_2D, &last_texture)
	//gl.GetIntegerv(gl.VIEWPORT, &last_viewport);
	gl.PushAttrib(gl.ENABLE_BIT | gl.COLOR_BUFFER_BIT | gl.TRANSFORM_BIT)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Disable(gl.CULL_FACE)
	gl.Disable(gl.DEPTH_TEST)
	gl.Enable(gl.SCISSOR_TEST)
	gl.EnableClientState(gl.VERTEX_ARRAY)
	gl.EnableClientState(gl.TEXTURE_COORD_ARRAY)
	gl.EnableClientState(gl.COLOR_ARRAY)
	gl.Enable(gl.TEXTURE_2D)
	//glUseProgram(0); // You may want this if using this code in an OpenGL 3+ context

	// Setup viewport, orthographic projection matrix
	gl.Viewport(0, 0, fb_width, fb_height)
	gl.MatrixMode(gl.PROJECTION)
	gl.PushMatrix()
	gl.LoadIdentity()
	gl.Ortho(0, io.DisplaySize.x, io.DisplaySize.y, 0, -1, 1)
	gl.MatrixMode(gl.MODELVIEW)
	gl.PushMatrix()
	gl.LoadIdentity()

	// Render command lists
	//#define OFFSETOF(TYPE, ELEMENT) ((size_t)&(((TYPE *)0)->ELEMENT))
	//cmdLists := (*drawData).GetCmdLists()
	for i := 0; i < (*drawData).GetCmdListsCount(); i++ {
		//cmdList := cmdLists.GetCmdBuffer()

		//vertexBuffer := cmdList.GetVtxBuffer()
		//indexBuffer := cmdList.GetIdxBuffer()

		//gl.VertexPointer(2, gl.FLOAT, sizeof(ImDrawVert), (void*)(vtx_buffer + OFFSETOF(ImDrawVert, pos)));
		//gl.TexCoordPointer(2, gl.FLOAT, sizeof(ImDrawVert), (void*)(vtx_buffer + OFFSETOF(ImDrawVert, uv)));
		//gl.ColorPointer(4, GL_UNSIGNED_BYTE, sizeof(ImDrawVert), (void*)(vtx_buffer + OFFSETOF(ImDrawVert, col)));

		//for cmd_i := 0; cmd_i < cmdList->CmdBuffer.size(); cmd_i++ {
		//	const ImDrawCmd* pcmd = &cmdList->CmdBuffer[cmd_i];
		//	if (pcmd->UserCallback) {
		//		pcmd->UserCallback(cmdList, pcmd);
		//	}
		//	else {
		//		gl.BindTexture(gl.TEXTURE_2D, );
		//		gl.Scissor((int)pcmd->ClipRect.x, (int)(fb_height - pcmd->ClipRect.w), (int)(pcmd->ClipRect.z - pcmd->ClipRect.x), (int)(pcmd->ClipRect.w - pcmd->ClipRect.y));
		//		gl.DrawElements(gl.TRIANGLES, (GLsizei)pcmd->ElemCount, sizeof(ImDrawIdx) == 2 ? GL_UNSIGNED_SHORT : GL_UNSIGNED_INT, idx_buffer);
		//	}
		//	idx_buffer += pcmd->ElemCount;
		//}
	}

	// Restore modified state
	gl.DisableClientState(gl.COLOR_ARRAY)
	gl.DisableClientState(gl.TEXTURE_COORD_ARRAY)
	gl.DisableClientState(gl.VERTEX_ARRAY)
	gl.BindTexture(gl.TEXTURE_2D, uint32(last_texture))
	gl.MatrixMode(gl.MODELVIEW)
	gl.PopMatrix()
	gl.MatrixMode(gl.PROJECTION)
	gl.PopMatrix()
	gl.PopAttrib()
	//gl.Viewport(last_viewport[0], last_viewport[1], (GLsizei)last_viewport[2], (GLsizei)last_viewport[3])

