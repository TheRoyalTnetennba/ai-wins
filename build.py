import glob
import shutil
import subprocess
import os

for filename in glob.iglob('/home/graham/go/src/ApiServer/*.go', recursive=True):
    os.unlink(filename)

for filename in glob.iglob('**/*.go', recursive=True):
    shutil.copy(filename, '/home/graham/go/src/ApiServer')

pkg_test = False

commands = ["go build ApiServer"]

if pkg_test:
    commands.insert(0, "git push")
    commands.insert(0, "git commit -am.")
    commands.insert(0, "git add -A")


for command in commands:
    processBuild = subprocess.Popen(command.split(), stdout=subprocess.PIPE)
    output, error = processBuild.communicate()
