// Copyright 2023 The Ebitengine Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build playstation5

// The actual implementation will be provided by -overlay.

#include "graphics_playstation5.h"

extern "C" int ebitengine_InitializeGraphics(void) {
  return 0;
}

extern "C" int ebitengine_NewImage(int* image, int width, int height) {
  return 0;
}

extern "C" int ebitengine_NewScreenFramebufferImage(int* image, int width, int height) {
  return 0;
}

extern "C" void ebitengine_DisposeImage(int id) {
}

extern "C" int ebitengine_NewShader(int* shader, const char* source) {
  return 0;
}

extern "C" void ebitengine_DisposeShader(int id) {
}
