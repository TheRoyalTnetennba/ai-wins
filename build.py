import glob
import shutil
import subprocess
import os

for filename in glob.iglob('/home/gpaye/go/src/ApiServer/*.go', recursive=True):
    os.unlink(filename)

for filename in glob.iglob('server/main/*.go', recursive=True):
    shutil.copy(filename, '/home/gpaye/go/src/ApiServer')

build = "go build ApiServer"
get_util = "go get -u github.com/TheRoyalTnetennba/ai-wins/server/utils"
get_tic_tac_toe = "go get -u github.com/TheRoyalTnetennba/ai-wins/server/games/tic-tac-toe"


processBuild = subprocess.Popen(build.split(), stdout=subprocess.PIPE)
output, error = processBuild.communicate()

processBuild = subprocess.Popen(get_util.split(), stdout=subprocess.PIPE)
output, error = processBuild.communicate()

processBuild = subprocess.Popen(
    get_tic_tac_toe.split(), stdout=subprocess.PIPE)
output, error = processBuild.communicate()
