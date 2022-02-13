from bs4 import BeautifulSoup as BS
import requests
import random
import json
from time import sleep
import os

#重启光猫
class Gateway:
    def __init__(self, username, password):
        self.base_url = "http://192.168.1.1"
        self.username = username
        self.password = password
        self.session = self.init_session()

    def init_session(self):
        session = requests.session()
        session.headers = {
            "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.80 Safari/537.36 Edg/98.0.1108.43"
        }
        login_url = f"{self.base_url}/admin/login.asp"
        r = session.get(login_url)
        html = BS(r.text, "html.parser")
        were_token = html.select_one('input[name="csrfmiddlewaretoken"]')
        csrf_token = html.select_one("input[name='csrftoken']")
        were_token = were_token.get("value")
        csrf_token = csrf_token.get("value")
        login_api = f"{self.base_url}/boaform/admin/formLogin"
        data = dict()
        data["username"] = self.username
        data["psd"] = self.password
        data["csrfmiddlewaretoken"] = were_token
        data["csrftoken"] = csrf_token
        data["username1"] = self.username
        data["username2"] = self.username
        data["psd1"] = self.password
        data["psd2"] = self.password
        session.post(login_api, data)
        return session

    def __restart_get_csf_token(self):
        api = f"{self.base_url}/mgm_dev_reboot.asp"
        r = self.session.get(api)
        html = BS(r.text, "html.parser")
        csf_token = html.select_one("input[name='csrftoken']")
        if not csf_token:
            self.init_session()
            # 递归
            self.__restart_get_csf_token()
        return csf_token.get("value")

    def restart_device(self):
        api = f"{self.base_url}/boaform/admin/formReboot"
        data = {
            "submit-url": "/mgm_dev_reboot.asp",
            "csrftoken": self.__restart_get_csf_token()
        }
        self.session.post(api, data)
        print("重启完成！")


def get_my_ip():
    url = "http://pv.sohu.com/cityjson"
    r = requests.get(url)
    r.encoding = "utf-8"
    result = r.text
    resultJson = result[19:-1]
    resultJson = json.loads(resultJson)
    return resultJson.get("cip")

#确保光猫重启后IP地址变了
def get_new_ip(username,psssword):
    while True:
        myIp = get_my_ip()
        # print(myIp)
        g = Gateway(username, psssword)
        g.restart_device()
        sleep(100)
        while True:
            try:
                requests.get("https://www.baidu.com", timeout=1)
                sleep(3)
                for i in range(10):
                    requests.get("https://www.baidu.com", timeout=1)
                break
            except:
                sleep(1)
                continue
        # newIp = get_my_ip()
        # print(newIp)
        if myIp == get_my_ip():
            print("ip没变")
        else:
            print("ip变")
            break



if __name__ == '__main__':
    count = 0
    myusername = ''
    mypassword = ''
    myurl = ''
    while count<20:
        session = requests.session()
        user_agent_list = [
            "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/22.0.1207.1 Safari/537.1"
            "Mozilla/5.0 (X11; CrOS i686 2268.111.0) AppleWebKit/536.11 (KHTML, like Gecko) Chrome/20.0.1132.57 Safari/536.11",
            "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1092.0 Safari/536.6",
            "Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1090.0 Safari/536.6",
            "Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/19.77.34.5 Safari/537.1",
            "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/536.5 (KHTML, like Gecko) Chrome/19.0.1084.9 Safari/536.5",
            "Mozilla/5.0 (Windows NT 6.0) AppleWebKit/536.5 (KHTML, like Gecko) Chrome/19.0.1084.36 Safari/536.5",
            "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1063.0 Safari/536.3",
            "Mozilla/5.0 (Windows NT 5.1) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1063.0 Safari/536.3",
            "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1063.0 Safari/536.3",
            "Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1062.0 Safari/536.3",
            "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1062.0 Safari/536.3",
            "Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.1 Safari/536.3",
            "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.1 Safari/536.3",
            "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.1 Safari/536.3",
            "Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
            "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.24 (KHTML, like Gecko) Chrome/19.0.1055.1 Safari/535.24",
            "Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/535.24 (KHTML, like Gecko) Chrome/19.0.1055.1 Safari/535.24"
        ]
        UserAgent = random.choice(user_agent_list)
        session.headers = UserAgent
        get_new_ip(myusername,mypassword)
        response = session.get(url=myurl, timeout=ord(os.urandom(1))%5+1)
        # print(response.status_code)
        if response.status_code == 200:
            count += 1
        else:
            sleep(30)
