import glob
import shutil
import subprocess
import os

for filename in glob.iglob('/home/gpaye/go/src/ApiServer/*.go', recursive=True):
    os.unlink(filename)

for filename in glob.iglob('server/main/*.go', recursive=True):
    shutil.copy(filename, '/home/gpaye/go/src/ApiServer')

commands = ["git add -A",
            "git commit -am.",
            "git push",
            "go get -u github.com/TheRoyalTnetennba/ai-wins/server/utils",
            "go get -u github.com/TheRoyalTnetennba/ai-wins/server/games/tic-tac-toe",
            "go build ApiServer"]

for command in commands:
    processBuild = subprocess.Popen(command.split(), stdout=subprocess.PIPE)
    output, error = processBuild.communicate()
