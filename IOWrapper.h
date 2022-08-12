#pragma once

#include "imguiWrapperTypes.h"

#ifdef __cplusplus
extern "C" {
#endif

extern IggBool iggWantCaptureMouse(IggIO handle);
extern IggBool iggWantCaptureKeyboard(IggIO handle);
extern IggBool iggWantTextInput(IggIO handle);

extern IggFontAtlas iggIoGetFonts(IggIO handle);

extern void iggIoSetDisplaySize(IggIO handle, IggVec2 const *value);
extern void iggIoSetMousePosition(IggIO handle, IggVec2 const *value);
extern void iggIoSetMouseButtonDown(IggIO handle, int index, IggBool value);
extern void iggIoAddMouseWheelDelta(IggIO handle, float x, float y);
extern void iggIoGetMouseWheelDelta(IggIO handle, float *vertical);
extern void iggIoGetMouseWheelHDelta(IggIO handle, float *horizontal);
extern void iggIoGetMouseDelta(IggIO handle, IggVec2 *delta);
extern void iggIoSetDeltaTime(IggIO handle, float value);
extern void iggIoSetFontGlobalScale(IggIO handle, float value);

extern IggBool iggIoGetMouseDrawCursor(IggIO handle);
extern void iggIoSetMouseDrawCursor(IggIO handle, IggBool value);

extern void iggIoAddKeyEvent(IggIO handle, int imguiKey, IggBool pressed);
extern void iggIoAddInputCharactersUTF8(IggIO handle, char const *utf8Chars);
extern void iggIoSetIniFilename(IggIO handle, char const *value);
extern void iggIoSetConfigFlags(IggIO handle, int flags);
extern int iggIoGetConfigFlags(IggIO handle);
extern void iggIoSetBackendFlags(IggIO handle, int flags);

extern void iggIoRegisterClipboardFunctions(IggIO handle);
extern void iggIoClearClipboardFunctions(IggIO handle);

extern int iggGetFrameCountSinceLastInput(IggIO handle);
extern void iggSetFrameCountSinceLastInput(IggIO handle, int count);

extern void iggIoAddFocusEvent(IggIO handle, IggBool focused);

#ifdef __cplusplus
}
#endif
