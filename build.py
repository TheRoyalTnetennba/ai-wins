import glob
import shutil
import subprocess
import os

for filename in glob.iglob('/home/gpaye/go/src/ApiServer/*.go', recursive=True):
    os.unlink(filename)

for filename in glob.iglob('server/**/*.go', recursive=True):
    shutil.copy(filename, '/home/gpaye/go/src/ApiServer')

bashCommand = "go build ApiServer"
process = subprocess.Popen(bashCommand.split(), stdout=subprocess.PIPE)
output, error = process.communicate()
