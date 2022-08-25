import subprocess

if __name__ == "__main__":
    temp = subprocess.Popen(["flatpak", "list"], stdout=subprocess.PIPE).stdout.read()
    temp2 = temp.splitlines()
    for i in range(len(temp2)):
        tm = temp2[i].split(b'\t')
        if len(tm) < 2:
            break
        print(tm[0].decode('UTF-8'), tm[1].decode('UTF-8'))