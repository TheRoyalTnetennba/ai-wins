import glob
import shutil
import subprocess
import os

for filename in glob.iglob('/home/gpaye/go/src/ApiServer/*.go', recursive=True):
    os.unlink(filename)

for filename in glob.iglob('server/main/*.go', recursive=True):
    shutil.copy(filename, '/home/gpaye/go/src/ApiServer')

build = "go build ApiServer"

processBuild = subprocess.Popen(build.split(), stdout=subprocess.PIPE)
output, error = processBuild.communicate()
