%module imgui
%{
#include "imgui.h"
%}

%rename(GetColorU32FromU32) GetColorU32(ImU32);
%rename(ValueUnsigned) Value(const char*, unsigned int);
%rename(Bool) ImGuiOnceUponAFrame::operator bool;
%rename(Get) ImGuiTextBuffer::operator[](int);
%rename(ImU32) ImColor::operator ImU32;
%rename(ImVec4) ImColor::operator ImVec4;
%rename(SetTexID_) ImFontAtlas::SetTexID(ImTextureID);
%rename(SetFallbackChar_) ImFont::SetFallbackChar(ImWchar);
%rename(ImGuiTextFilter_TextRange) ImGuiTextFilter::TextRange;
%rename(ImGuiStorage_Pair) ImGuiStorage::Pair;
%rename(ImFontAtlas_GlyphRangesBuilder) ImFontAtlas::GlyphRangesBuilder;
%rename(ImFontAtlas_CustomRect) ImFontAtlas::CustomRect;
%rename(ImFont_Glyph) ImFont::Glyph;

%include "imgui.h"
