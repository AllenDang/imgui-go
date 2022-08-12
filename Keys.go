package imgui

// Constants to fill IO.KeyMap() lookup with indices into the IO.KeysDown[512] array.
// The mapped indices are then the ones reported to IO.KeyPress() and IO.KeyRelease().
const (
	// Keyboard
	ImGuiKey_None = 0

	ImGuiKey_Tab = iota + 512 // == ImGuiKey_NamedKey_BEGIN
	ImGuiKey_LeftArrow

	ImGuiKey_RightArrow

	ImGuiKey_UpArrow

	ImGuiKey_DownArrow

	ImGuiKey_PageUp

	ImGuiKey_PageDown

	ImGuiKey_Home

	ImGuiKey_End

	ImGuiKey_Insert

	ImGuiKey_Delete

	ImGuiKey_Backspace

	ImGuiKey_Space

	ImGuiKey_Enter

	ImGuiKey_Escape

	ImGuiKey_LeftCtrl
	ImGuiKey_LeftShift
	ImGuiKey_LeftAlt
	ImGuiKey_LeftSuper

	ImGuiKey_RightCtrl
	ImGuiKey_RightShift
	ImGuiKey_RightAlt
	ImGuiKey_RightSuper

	ImGuiKey_Menu

	ImGuiKey_0
	ImGuiKey_1
	ImGuiKey_2
	ImGuiKey_3
	ImGuiKey_4
	ImGuiKey_5
	ImGuiKey_6
	ImGuiKey_7
	ImGuiKey_8
	ImGuiKey_9

	ImGuiKey_A
	ImGuiKey_B
	ImGuiKey_C
	ImGuiKey_D
	ImGuiKey_E
	ImGuiKey_F
	ImGuiKey_G
	ImGuiKey_H
	ImGuiKey_I
	ImGuiKey_J

	ImGuiKey_K
	ImGuiKey_L
	ImGuiKey_M
	ImGuiKey_N
	ImGuiKey_O
	ImGuiKey_P
	ImGuiKey_Q
	ImGuiKey_R
	ImGuiKey_S
	ImGuiKey_T

	ImGuiKey_U
	ImGuiKey_V
	ImGuiKey_W
	ImGuiKey_X
	ImGuiKey_Y
	ImGuiKey_Z

	ImGuiKey_F1
	ImGuiKey_F2
	ImGuiKey_F3
	ImGuiKey_F4
	ImGuiKey_F5
	ImGuiKey_F6

	ImGuiKey_F7
	ImGuiKey_F8
	ImGuiKey_F9
	ImGuiKey_F10
	ImGuiKey_F11
	ImGuiKey_F12

	ImGuiKey_Apostrophe

	ImGuiKey_Comma

	ImGuiKey_Minus

	ImGuiKey_Period

	ImGuiKey_Slash

	ImGuiKey_Semicolon

	ImGuiKey_Equal

	ImGuiKey_LeftBracket

	ImGuiKey_Backslash

	ImGuiKey_RightBracket

	ImGuiKey_GraveAccent

	ImGuiKey_CapsLock

	ImGuiKey_ScrollLock

	ImGuiKey_NumLock

	ImGuiKey_PrintScreen

	ImGuiKey_Pause

	ImGuiKey_Keypad0
	ImGuiKey_Keypad1
	ImGuiKey_Keypad2
	ImGuiKey_Keypad3
	ImGuiKey_Keypad4

	ImGuiKey_Keypad5
	ImGuiKey_Keypad6
	ImGuiKey_Keypad7
	ImGuiKey_Keypad8
	ImGuiKey_Keypad9

	ImGuiKey_KeypadDecimal

	ImGuiKey_KeypadDivide

	ImGuiKey_KeypadMultiply

	ImGuiKey_KeypadSubtract

	ImGuiKey_KeypadAdd

	ImGuiKey_KeypadEnter

	ImGuiKey_KeypadEqual

	ImGuiKey_GamepadStart

	ImGuiKey_GamepadBack

	ImGuiKey_GamepadFaceUp

	ImGuiKey_GamepadFaceDown

	ImGuiKey_GamepadFaceLeft

	ImGuiKey_GamepadFaceRight

	ImGuiKey_GamepadDpadUp

	ImGuiKey_GamepadDpadDown

	ImGuiKey_GamepadDpadLeft

	ImGuiKey_GamepadDpadRight
	ImGuiKey_GamepadL1

	ImGuiKey_GamepadR1

	ImGuiKey_GamepadL2

	ImGuiKey_GamepadR2

	ImGuiKey_GamepadL3

	ImGuiKey_GamepadR3

	ImGuiKey_GamepadLStickUp

	ImGuiKey_GamepadLStickDown

	ImGuiKey_GamepadLStickLeft

	ImGuiKey_GamepadLStickRight

	ImGuiKey_GamepadRStickUp

	ImGuiKey_GamepadRStickDown

	ImGuiKey_GamepadRStickLeft

	ImGuiKey_GamepadRStickRight

	ImGuiKey_ModCtrl
	ImGuiKey_ModShift
	ImGuiKey_ModAlt
	ImGuiKey_ModSuper

	ImGuiKey_COUNT
)
