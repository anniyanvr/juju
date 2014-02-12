; Script generated by the Inno Setup Script Wizard.
; SEE THE DOCUMENTATION FOR DETAILS ON CREATING INNO SETUP SCRIPT FILES!

#define MyAppName "Juju"
#define MyAppVersion "1.16.7"
#define MyAppPublisher "Canonical, Ltd"
#define MyAppURL "http://juju.ubuntu.com/"
#define MyAppExeName "juju.exe"

[Setup]
; NOTE: The value of AppId uniquely identifies this application.
; Do not use the same AppId value in installers for other applications.
; (To generate a new GUID, click Tools | Generate GUID inside the IDE.)
AppId={{B2781001-AA89-4E70-AC2B-17D004DFBAE2}
AppName={#MyAppName}
AppVersion={#MyAppVersion}
;AppVerName={#MyAppName} {#MyAppVersion}
AppPublisher={#MyAppPublisher}
AppPublisherURL={#MyAppURL}
ChangesEnvironment=yes
DefaultDirName={pf}\{#MyAppName}
DisableDirPage=auto
DefaultGroupName={#MyAppName}
DisableProgramGroupPage=yes
OutputBaseFilename=juju-setup-{#MyAppVersion}
SetupIconFile=juju.ico
Compression=lzma
SolidCompression=yes
UninstallDisplayIcon={app}\juju.ico
UninstallDisplayName=Juju
WizardImageStretch=no
WizardImageFile=juju-wizard-side.bmp
WizardImageBackColor=clWhite
WizardSmallImageFile=juju-55.bmp
LicenseFile=LICENCE.txt

[Languages]
Name: "english"; MessagesFile: "compiler:Default.isl"

[Tasks]
Name: modifypath; Description: Add application directory to your environment path

[Files]
Source: "juju.exe"; DestDir: "{app}"; Flags: ignoreversion
Source: "juju.ico"; DestDir: "{app}"; Flags: ignoreversion
Source: "README.txt"; DestDir: "{app}"; Flags: ignoreversion
Source: "LICENCE.txt"; DestDir: "{app}"; Flags: ignoreversion

[Run]
Filename: "{app}\README.txt"; Description: "View the README file"; Flags: postinstall shellexec skipifsilent

[Code]
const
	ModPathName= 'modifypath';
  ModPathType = 'system';

function ModPathDir(): TArrayOfString;
begin
	setArrayLength(Result, 1)
	Result[0] := ExpandConstant('{app}');
end;

#include "modpath.iss"
