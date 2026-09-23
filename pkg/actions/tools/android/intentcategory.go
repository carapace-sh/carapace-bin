package android

import (
	"github.com/carapace-sh/carapace"
)

// ActionIntentCategories completes intent categories
//
//	android.intent.category.DEFAULT (Set if the activity should be an option for the default action (center press) to perform on a piece of data.)
//	android.intent.category.BROWSABLE (Activities that can be safely invoked from a browser must support this category.)
//	android.intent.category.VOICE (Categories for activities that can participate in voice interaction.)
func ActionIntentCategories() carapace.Action {
	return carapace.ActionValuesDescribed(
		"android.intent.category.DEFAULT", "Set if the activity should be an option for the default action (center press) to perform on a piece of data.",
		"android.intent.category.BROWSABLE", "Activities that can be safely invoked from a browser must support this category.",
		"android.intent.category.VOICE", "Categories for activities that can participate in voice interaction.",
		"android.intent.category.ALTERNATIVE", "Set if the activity should be considered as an alternative action to the data the user is currently viewing.",
		"android.intent.category.SELECTED_ALTERNATIVE", "Set if the activity should be considered as an alternative selection action to the data the user has currently selected.",
		"android.intent.category.TAB", "Intended to be used as a tab inside of a containing TabActivity",
		"android.intent.category.LAUNCHER", "Should be displayed in the top-level launcher",
		"android.intent.category.LEANBACK_LAUNCHER", "Indicates an activity optimized for Leanback mode, and that should be displayed in the Leanback launcher",
		"android.intent.category.CAR_LAUNCHER", "Indicates the preferred entry-point activity when an application is launched from a Car launcher.",
		"android.intent.category.COMMUNAL_MODE", "Used to indicate that the activity can be used in communal mode",
		"android.intent.category.LEANBACK_SETTINGS", "Indicates a Leanback settings activity to be displayed in the Leanback launcher",
		"android.intent.category.INFO", "Provides information about the package it is in; typically used if a package does not contain a",
		"android.intent.category.HOME", "This is the home activity, that is the first activity that is displayed when the device boots",
		"android.intent.category.HOME_MAIN", "This is the home activity that is displayed when the device is finished setting up and ready for use",
		"android.intent.category.SECONDARY_HOME", "The home activity shown on secondary displays that support showing home activities",
		"android.intent.category.SETUP_WIZARD", "This is the setup wizard activity, that is the first activity that is displayed when the user sets up the device for the first time",
		"android.intent.category.LAUNCHER_APP", "This is the home activity, that is the activity that serves as the launcher app from there the user can start other apps.",
		"android.intent.category.PREFERENCE", "This activity is a preference panel",
		"android.intent.category.DEVELOPMENT_PREFERENCE", "This activity is a development preference panel",
		"android.intent.category.EMBED", "Capable of running inside a parent activity container",
		"android.intent.category.APP_MARKET", "This activity allows the user to browse and download new applications",
		"android.intent.category.MONKEY", "This activity may be exercised by the monkey or other automated test tools",
		"android.intent.category.TEST", "To be used as a test (not part of the normal user experience)",
		"android.intent.category.UNIT_TEST", "To be used as a unit test (run through the Test Harness)",
		"android.intent.category.SAMPLE_CODE", "To be used as a sample code example (not part of the normal user experience)",
		"android.intent.category.OPENABLE", "Used to indicate that an intent only wants URIs that can be opened with",
		"android.intent.category.CAR_DOCK", "An activity to run when device is inserted into a car dock.",
		"android.intent.category.DESK_DOCK", "An activity to run when device is inserted into a desk dock.",
		"android.intent.category.LE_DESK_DOCK", "An activity to run when device is inserted into a analog (low end) dock.",
		"android.intent.category.HE_DESK_DOCK", "An activity to run when device is inserted into a digital (high end) dock.",
		"android.intent.category.CAR_MODE", "Used to indicate that the activity can be used in a car environment",
		"android.intent.category.VR_HOME", "An activity to use for the launcher when the device is placed in a VR Headset viewer.",
		"android.intent.category.APP_BROWSER", "Used with",
		"android.intent.category.APP_CALCULATOR", "Used with",
		"android.intent.category.APP_CALENDAR", "Used with",
		"android.intent.category.APP_CONTACTS", "Used with",
		"android.intent.category.APP_EMAIL", "Used with",
		"android.intent.category.APP_GALLERY", "Used with",
		"android.intent.category.APP_MAPS", "Used with",
		"android.intent.category.APP_MESSAGING", "Used with",
		"android.intent.category.APP_MUSIC", "Used with",
		"android.intent.category.APP_FILES", "Used with",
		"android.intent.category.APP_WEATHER", "Used with",
		"android.intent.category.APP_FITNESS", "Used with",
	)
}
