#!/bin/sh
sed -i -e 's/ImGui::TreePush..;/ImGui::TreePush((char *)0);/g' imgui_wrap.cxx
echo 'Patched imgui_wrap.cxx.'
