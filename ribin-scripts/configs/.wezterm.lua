-- Pull in the wezterm API
local wezterm = require("wezterm")

-- This table will hold the configuration.
local config = {}

-- In newer versions of wezterm, use the config_builder which will
-- help provide clearer error messages
if wezterm.config_builder then
	config = wezterm.config_builder()
end

-- This is where you actually apply your config choices

-- For example, changing the color scheme:
-- scheme_list: https://wezfurlong.org/wezterm/colorschemes/a/index.html#adventuretime
-- config.color_scheme = "Low Contrast (terminal.sexy)"

config.background = {
	{
		source = {
			File = "/Users/ribincao/Downloads/logo/bud_master.png",
		},
		-- The texture tiles vertically but not horizontally.
		-- When we repeat it, mirror it so that it appears "more seamless".
		-- An alternative to this is to set `width = "100%"` and have
		-- it stretch across the display
		-- repeat_y = "Repeat",
		width = "100%",
		hsb = { brightness = 0.382 },
		-- When the viewport scrolls, move this layer 10% of the number of
		-- pixels moved by the main viewport. This makes it appear to be
		-- further behind the text.
		-- attachment = { Parallax = 0.1 },
	},
}
-- config.colors = {
-- 	background = "purple",
-- }

-- font_list: wezterm ls-fonts --list-system
config.font = wezterm.font("JetBrains Mono", { weight = "ExtraLight", stretch = "Normal", style = "Normal" }) -- <built-in>, BuiltIn
-- config.font = wezterm.font("JetBrains Mono", { weight = "Bold", stretch = "Normal", style = "Normal" }) -- <built-in>, BuiltIn
-- wezterm.font("JetBrains Mono", { weight = "Bold", italic = false })
-- config.font = wezterm.font("GoMono NF", { weight = "Bold", italic = false })

-- and finally, return the configuration to wezterm
return config
