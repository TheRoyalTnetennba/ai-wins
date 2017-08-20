import glob
import shutil
import subprocess

for filename in glob.iglob('server/**/*.go', recursive=True):
    shutil.copy(filename, '/home/gpaye/go/src/ApiServer')

bashCommand = "go build ApiServer"
process = subprocess.Popen(bashCommand.split(), stdout=subprocess.PIPE)
output, error = process.communicate()
