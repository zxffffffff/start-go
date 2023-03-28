SET pbExe=%~dp0\protoc-win64\bin\protoc.exe
SET pbDir=%~dp0\pb
SET goDir=%~dp0\pb_go

if exist %goDir% ( rd /s /Q %goDir% )
md %goDir%

%pbExe% ^
--proto_path=%pbDir% ^
--go_out=%goDir% ^
%pbDir%\common\*.proto ^
%pbDir%\notify\*.proto ^
%pbDir%\req\*.proto ^
%pbDir%\res\*.proto ^
%pbDir%\*.proto
