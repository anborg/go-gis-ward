call build.bat
SET BASE_PATH=\\hello\test
xcopy  /f   /y %BINARY_NAME%   %BASE_PATH%\
xcopy  /f   /y config-mkdev.yml   %BASE_PATH%\
xcopy  /f /v  /i /y assets   %BASE_PATH%\assets
xcopy  /f /v  /i /y templates   %BASE_PATH%\templates
:: %BASE_PATH%\%BINARY_NAME% -configFile=%BASE_PATH%\config1.yml