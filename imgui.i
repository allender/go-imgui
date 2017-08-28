%module imgui
%{
#include "imgui.h"
%}

%rename("%(regex:/^(?:Im|ImGui|Gui)(.*)_?$/\\1/)s") "";
%rename("%s", %$isnested) "";
%rename(GetColorU32FromU32) GetColorU32(ImU32);
%rename(ValueUnsigned) Value(const char*, unsigned int);
%rename(Bool) ImGuiOnceUponAFrame::operator bool;
%rename(Get) ImGuiTextBuffer::operator[](int);
%rename(U32) ImColor::operator ImU32;
%rename(Vec4) ImColor::operator ImVec4;
%rename(SetTexID_) ImFontAtlas::SetTexID(ImTextureID);
%rename(SetFallbackChar_) ImFont::SetFallbackChar(ImWchar);

%include "imgui.h"
/*
%ignore ImVector::swap;

%template(FontConfigList) ImVector<ImFontConfig>;
%template(Float32List) ImVector<float>;
%template(ByteList) ImVector<char>;
%template(DrawCmdList) ImVector<ImDrawCmd>;
%template(DrawIdxList) ImVector<ImDrawIdx>;
%template(DrawVertList) ImVector<ImDrawVert>;
%template(Vec4List) ImVector<ImVec4>;
%template(TextureIDList) ImVector<ImTextureID>;
%template(Vec2List) ImVector<ImVec2>;
%template(DrawChannelList) ImVector<ImDrawChannel>;
*/
%typemap(gotype) (ImDrawList** CmdLists) %{[]DrawList%};
%typemap(gotype) (long) %{int64%};
%typemap(gotype) (unsigned long) %{uint64%};
