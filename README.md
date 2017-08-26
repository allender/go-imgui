# go-imgui
Go bindings for the awesome [dear imgui](https://github.com/ocornut/imgui) immediate mode GUI library, written in C++.
These bindings are generated with SWIG.

**Current imgui version**: 1.51

# Installation
Run this command to update or install the library:
```
go get -u https://github.com/armored-dragon/go-imgui
```

# Usage
Usage is more or less identical to the C++ library.
For example, `imgui.ShowTestWindow()` will show the built-in demo window, which contains usage of just about every feature imgui has to offer.
All `Im` type prefixes are removed, so types like `ImVec2` or `ImFontAtlas` become `imgui.Vec2` and `imgui.FontAtlas`.
`ImGui` type prefixes are also removed.
Prefixes of flag constants are also removed, so `ImGuiWindowFlags_NoResize` will become `imgui.WindowFlags_NoResize`, and so on for other flags.

If you import the package under a different name, just use `packagename.` instead of `imgui.`.

### Nested Structs
Certain nested structs present in imgui, such as Glyph (`ImFont::Glyph`), are not available in these bindings.
This is because SWIG doesn't support nested structs.
Full list of such structs:
  - `ImGuiTextFilter::TextRange`
  - `ImGuiStorage::Pair`
  - `ImFontAtlas::GlyphRangesBuilder`
  - `ImFontAtlas::CustomRect`
  - `ImFont::Glyph`

# Generating
To generate the bindings with SWIG, this command is used:
```
swig -go -cgo -c++ -intgosize 32 imgui.i
```

**Note**: The result will not compile unless `patch.sh` is run. This is because SWIG fails to recognize a case of function overload ambiguity and the resulting call fails to compile.
