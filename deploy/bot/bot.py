from selenium import webdriver
import time


def visit(url):
    with open("flag", "r") as f:
        flag = f.read()
        
    try:
        driver.get(url)
        driver.add_cookie({"name": "flag", "value": flag})
        driver.get(url)
        time.sleep(2)
    except:
        pass
    finally:
        driver.quit()
        
        

if __name__ == "__main__":
    options = webdriver.ChromeOptions()
    options.add_argument('--headless')
    options.add_argument('--no-sandbox')
    options.add_argument('--disable-dev-shm-usage')
    options.add_argument('--remote-debugging-port=9222')
    options.binary_location = "/usr/bin/chromium"
    
    driver = webdriver.Chrome(options=options)
    url = input("input url> ")
    visit(url)
