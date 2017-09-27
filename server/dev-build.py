import glob
import shutil
import subprocess
import os

for filename in glob.iglob('/home/gpaye8/go/src/ApiWins/*.go', recursive=True):
    os.unlink(filename)

for filename in glob.iglob('server/main/*.go', recursive=True):
    shutil.copy(filename, '/home/gpaye8/go/src/ApiWins')

pkg_test = False

commands = ["go get -u github.com/TheRoyalTnetennba/ai-wins/server/utils",
            "go get -u github.com/TheRoyalTnetennba/ai-wins/server/ai",
            "go build ApiWins"]

if pkg_test:
    commands.insert(0, "git push")
    commands.insert(0, "git commit -am.")
    commands.insert(0, "git add -A")


for command in commands:
    processBuild = subprocess.Popen(command.split(), stdout=subprocess.PIPE)
    output, error = processBuild.communicate()
