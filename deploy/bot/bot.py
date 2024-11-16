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
    
    driver = webdriver.Chrome(options=options)
    url = input("input url> ")
    visit(url)
