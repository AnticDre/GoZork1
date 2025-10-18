@echo off
set /p REPO_NAME="Enter repository name: "
set /p USERNAME="Enter GitHub username: "
set /p COMMIT_MSG="Enter commit message (or press Enter for 'Initial commit'): "

if "%COMMIT_MSG%"=="" set COMMIT_MSG=Initial commit

echo Setting up Git repository: %REPO_NAME%

git init
git add .
git commit -m "%COMMIT_MSG%"
git branch -M main
git remote add origin https://github.com/%USERNAME%/%REPO_NAME%.git

echo.
echo Repository setup complete!
echo Remote URL: https://github.com/%USERNAME%/%REPO_NAME%.git
echo Next step: Create the repository on GitHub, then run:
echo git push -u origin main