<?php
# See includes/DefaultSettings.php for all configurable settings
# and their default values, but don't forget to make changes in _this_
# file, not there.
#
# Further documentation for configuration settings may be found at:
# https://www.mediawiki.org/wiki/Manual:Configuration_settings

# Protect against web entry
if ( !defined( 'MEDIAWIKI' ) ) {
	exit;
}

$wgReadOnly = getenv('MW_WG_READONLY') ?: false;
ini_set('display_errors', false);

if (getenv('MW_DEBUG')) {
    $wgShowExceptionDetails = true;
    $wgShowDBErrorBacktrace = true;
    $wgDebugToolbar = true;
}

## Uncomment this to disable output compression
# $wgDisableOutputCompression = true;

$wgSitename = getenv('MW_SITENAME');
$wgMetaNamespace = getenv('MW_METANAMESPACE');

## The URL base path to the directory containing the wiki;
## defaults for all runtime URL paths are based off of this.
## For more information on customizing the URLs
## (like /w/index.php/Page_title to /wiki/Page_title) please see:
## https://www.mediawiki.org/wiki/Manual:Short_URL
$wgScriptPath = "";
$wgScriptExtension = ".php";

## Path to articles is set up so that pages are reachable on /Page_Name
$wgArticlePath = "/$1";

## The protocol and server name to use in fully-qualified URLs
$wgServer = getenv('MW_SERVER');

## The relative URL path to the skins directory
$wgStylePath = "$wgScriptPath/skins";

## The relative URL path to the logo.  Make sure you change this from the default,
## or else you'll overwrite your logo when you upgrade!
$wgLogo = getenv("MW_LOGO") ?: "$wgScriptPath/resources/assets/wiki.png";

## UPO means: this is also a user preference option

$wgEnableEmail = getenv("MW_ENABLE_EMAIL") ?: false;
$wgEnableUserEmail = true; # UPO

$wgEmergencyContact = getenv("MW_EMERGENCY_CONTACT") ?: "apache@localhost";
$wgPasswordSender = getenv("MW_PASSWORD_SENDER") ?: "apache@localhost";

$wgEnotifUserTalk = false; # UPO
$wgEnotifWatchlist = false; # UPO
$wgEmailAuthentication = true;

## Database settings
$wgDBtype = "mysql";
$wgDBserver = "mysql";
$wgDBname = "mediawiki";
$wgDBuser = getenv('MW_DBUSER');
$wgDBpassword = getenv('MW_DBPASS');

# MySQL specific settings
$wgDBprefix = "";

# MySQL table options to use during installation or update
$wgDBTableOptions = "ENGINE=InnoDB, DEFAULT CHARSET=utf8";

# Experimental charset support for MySQL 5.0.
$wgDBmysql5 = true;

## Shared memory settings
$wgMainCacheType = CACHE_NONE;
$wgMemCachedServers = array();

## To enable image uploads, make sure the 'images' directory
## is writable, then set this to true:
$wgEnableUploads = true;
$wgUseImageMagick = true;
$wgImageMagickConvertCommand = "/usr/bin/convert";

# InstantCommons allows wiki to use images from http://commons.wikimedia.org
$wgUseInstantCommons = false;

## If you use ImageMagick (or any other shell command) on a
## Linux server, this will need to be set to the name of an
## available UTF-8 locale
$wgShellLocale = getenv("MW_SHELL_LOCALE") ?: "C.UTF-8";

## If you want to use image uploads under safe mode,
## create the directories images/archive, images/thumb and
## images/temp, and make them all writable. Then uncomment
## this, if it's not already uncommented:
#$wgHashedUploadDirectory = false;

## Set $wgCacheDirectory to a writable directory on the web server
## to make your wiki go slightly faster. The directory should not
## be publically accessible from the web.
#$wgCacheDirectory = "$IP/cache";

# Site language code, should be one of the list in ./languages/Names.php
$wgLanguageCode = getenv("MW_LANGUAGECODE");

$wgSecretKey = getenv("MW_SECRET");

# Site upgrade key. Must be set to a string (default provided) to turn on the
# web installer while LocalSettings.php is in place
$wgUpgradeKey = getenv("MW_UPGRADE_KEY") ?: "b31022590a7b3b8f";

if (getenv('MW_DISABLE_API')) {
    $wgEnableAPI = false;
}
if (getenv('MW_DISABLE_FEED')) {
    $wgFeed = false;
}

if (getenv('MW_REFERRER_POLICY')) {
    $wgReferrerPolicy = getenv('MW_REFERRER_POLICY');
}

## For attaching licensing metadata to pages, and displaying an
## appropriate copyright notice / icon. GNU Free Documentation
## License and Creative Commons licenses are supported so far.
$wgRightsPage = ""; # Set to the title of a wiki page that describes your license/copyright
$wgRightsUrl = "";
$wgRightsText = "";
$wgRightsIcon = "";

# Path to the GNU diff3 utility. Used for conflict resolution.
$wgDiff3 = "/usr/bin/diff3";

## Default skin: you can change the default skin. Use the internal symbolic
## names, ie 'vector',
$wgDefaultSkin = "vector";

# Enabled skins.
# The following skins were automatically enabled:
wfLoadSkin( 'Vector' );

# Subpages
if (getenv('MW_NS_WITH_SUBPAGES_MAIN')) {
    $wgNamespacesWithSubpages[NS_MAIN] = true;
}
if (getenv('MW_NS_WITH_SUBPAGES_TEMPLATE')) {
    $wgNamespacesWithSubpages[NS_TEMPLATE] = true;
}

wfLoadExtension( 'Interwiki' );
wfLoadExtension( 'Renameuser' );
wfLoadExtension( 'ParserFunctions' );
if (getenv('MV_PARSERFUNCTIONS_ENABLE_STRING_FUNCTIONS')) {
    $wgPFEnableStringFunctions = true;
}

# VisualEditor Extension
wfLoadExtension( 'VisualEditor' );

# Enable by default for everybody
$wgDefaultUserOptions['visualeditor-enable'] = 1;

# Semantic Stuff
require_once "$IP/extensions/SemanticMediaWiki/SemanticMediaWiki.php";
enableSemantics();
require_once "$IP/extensions/SemanticInternalObjects/SemanticInternalObjects.php";
require_once "$IP/extensions/SemanticInternalObjects/SemanticInternalObjects.php";
require_once "$IP/extensions/SemanticCompoundQueries/SemanticCompoundQueries.php";
require_once "$IP/extensions/SemanticDrilldown/SemanticDrilldown.php";
wfLoadExtension( 'PageForms' );

wfLoadExtension( 'Arrays' );

if (getenv('MW_RAWHTML') === 'true') {
    $wgRawHtml = true;
}
