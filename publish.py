import pwd
import subprocess
import time
from pathlib import Path
import zipfile
import os
import oss2
import shutil

def path_is_parent(parent_path, child_path):
    parent_path = os.path.abspath(parent_path)
    child_path = os.path.abspath(child_path)
    return os.path.commonpath([parent_path]) == os.path.commonpath([parent_path, child_path])

def is_ignored(target_path, ignores):
    for ignore in ignores:
        if os.path.relpath(target_path, ignore) == '.' or path_is_parent(ignore, target_path):
            return True
    return False

def zip_file(workspace, eve_app):
    os.chdir('%s/%s/src' % (workspace, eve_app))
    if os.path.isfile('README.md'):
        shutil.copy('README.md', 'code')
    elif os.path.isfile('readme.md'):
        shutil.copy('readme.md', 'code')
    elif os.path.isfile('Readme.md'):
        shutil.copy('Readme.md', 'code')
    os.chdir('%s/%s/src/code' % (workspace, eve_app))
    ignore_list = ['./.git', './.github', './.idea', './.DS_Store', './.vscode']
    with zipfile.ZipFile('code.zip', mode="w") as f:
        for dirpath, dirnames, filenames in os.walk('./'):
            if dirpath != './' and is_ignored(dirpath, ignore_list):
                continue
            for filename in filenames:
                absoult_file_path = os.path.join(dirpath, filename)
                if not is_ignored(absoult_file_path, ignore_list) and "code.zip" not in filename:
                    f.write(os.path.join(dirpath, filename))

    return 'code.zip'

def zip_golang_binary(workspace, eve_app):
    os.chdir('%s/%s/src' % (workspace, eve_app))
    if os.path.isfile('README.md'):
        shutil.copy('README.md', 'code/target')
    elif os.path.isfile('readme.md'):
        shutil.copy('readme.md', 'code/target')
    elif os.path.isfile('Readme.md'):
        shutil.copy('Readme.md', 'code/target')
    os.chdir('%s/%s/src/code/target' % (workspace, eve_app))
    ignore_list = ['./.git', './.github', './.idea', './.DS_Store', './.vscode']
    with zipfile.ZipFile('code.zip', mode="w") as f:
        for dirpath, dirnames, filenames in os.walk('./'):
            if dirpath != './' and is_ignored(dirpath, ignore_list):
                continue
            for filename in filenames:
                absoult_file_path = os.path.join(dirpath, filename)
                if not is_ignored(absoult_file_path, ignore_list) and "code.zip" not in filename:
                    f.write(os.path.join(dirpath, filename))

    return 'code.zip'

def upload_oss(code_name, zip_file):
    auth = oss2.Auth(os.environ.get('AccessKeyId'), os.environ.get('AccessKeySecret'))
    bucket = oss2.Bucket(auth, os.environ.get('ArtifactEndpoint'), os.environ.get('ArtifactBucket'))

    with open(zip_file, 'rb') as fileobj:
        object_name = '%s/code.zip' % (code_name)
        bucket.put_object(object_name, fileobj)

workspace = os.getcwd()
with open('update.list') as f:
    publish_list = [eve_app.strip() for eve_app in f.readlines()]

failed_registry = []
failed_oss = []
for eve_app in publish_list:
    times = 1
    os.chdir(workspace)
    try:
        while times <= 3:
            print("----------------------: ", eve_app)
            # publish app to registry 
            publish_script = 'https://serverless-registry.oss-cn-hangzhou.aliyuncs.com/publish-file/python3/hub-publish.py'
            command = 'cd %s && wget %s && python hub-publish.py' % (
                eve_app, publish_script)
            child = subprocess.Popen(
                command, stdout=subprocess.PIPE, stderr=subprocess.PIPE, shell=True, )
            stdout, stderr = child.communicate()
            if child.returncode == 0:
                print("stdout:", stdout.decode("utf-8"))
                break
            else:
                print("stdout:", stdout.decode("utf-8"))
                print("stderr:", stderr.decode("utf-8"))
                time.sleep(3)
                if times == 3:
                    raise ChildProcessError(stderr)
            times = times + 1
    except Exception as e:
        print('Failed to publish registry, app %s, err: %s' % (eve_app, e))
        failed_registry.append(eve_app)

    # publish code.zip to oss
    os.chdir(workspace)
    try:
        makefile = Path('%s/src/Makefile' % (eve_app))
        if makefile.is_file():
            print("----------------------Makefile: ", makefile)
            command = 'cd %s/src && make release' % (eve_app)
            print(command)
            child = subprocess.Popen(
                command, stdout=subprocess.PIPE, stderr=subprocess.PIPE, shell=True, )
            stdout, stderr = child.communicate()
            print("stdout:", stdout.decode("utf-8"))
            if child.returncode != 0:
                print("stderr:", stderr.decode("utf-8"))
                raise ChildProcessError(stderr)
        jarPath = '%s/%s/src/code/target/code.jar' % (workspace, eve_app)
        golangBinaryPath = '%s/%s/src/code/target/main' % (workspace, eve_app)
        if os.path.isfile(jarPath):
            code_zip = jarPath   
        elif os.path.isfile(golangBinaryPath):
            code_zip = zip_golang_binary(workspace, eve_app)
        else:
            code_zip = zip_file(workspace, eve_app)
        upload_oss(eve_app, code_zip)
    except Exception as e:
        print('Failed to publish oss, app %s, err: %s' % (eve_app, e)) 
        failed_oss.append(eve_app)
    
print('Failed registry list: ', failed_registry)
print('Failed oss list: ', failed_oss)