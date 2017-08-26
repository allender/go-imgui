%module imgui
%{
#include "imgui.h"
%}

%rename("strip:[Im]") "";
%rename(GetColorU32FromU32) GetColorU32(ImU32);
%rename(ValueUnsigned) Value(const char*, unsigned int);
%rename(Bool) ImGuiOnceUponAFrame::operator bool;
%rename(Get) ImGuiTextBuffer::operator[](int);
%rename(U32) ImColor::operator ImU32;
%rename(Vec4) ImColor::operator ImVec4;
%rename(SetTexID_) ImFontAtlas::SetTexID(ImTextureID);
%rename(SetFallbackChar_) ImFont::SetFallbackChar(ImWchar);
%rename(TextRange) ImGuiTextFilter::TextRange;
%rename(Pair) ImGuiStorage::Pair;
%rename(GlyphRangesBuilder) ImFontAtlas::GlyphRangesBuilder;
%rename(CustomRect) ImFontAtlas::CustomRect;
%rename(Glyph) ImFont::Glyph;

%include "imgui.h"
