import glob
import shutil
import subprocess
import os

for filename in glob.iglob('/home/graham/go/src/ApiServer/*.go', recursive=True):
    os.unlink(filename)

for filename in glob.iglob('server/main/*.go', recursive=True):
    shutil.copy(filename, '/home/graham/go/src/ApiServer')

pkg_test = False

commands = ["go get -u github.com/TheRoyalTnetennba/ai-wins/server/utils",
            "go get -u github.com/TheRoyalTnetennba/ai-wins/server/games/tic-tac-toe",
            "go build ApiServer"]

if pkg_test:
    commmands.insert(0, "git push")
    commmands.insert(0, "git commit -am.")
    commmands.insert(0, "git add -A")


for command in commands:
    processBuild = subprocess.Popen(command.split(), stdout=subprocess.PIPE)
    output, error = processBuild.communicate()
