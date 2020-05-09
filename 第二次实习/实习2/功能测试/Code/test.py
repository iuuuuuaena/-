# coding:utf-8
# @Project Name spider
# @Name test.py
# @Author by jxd
# @Email waniiu@126.com
# @Time 2020/4/28 3:40 下午
# @Method for selenium
from selenium import webdriver
from selenium.webdriver.common.by import By  # 定位器
from selenium.webdriver.common.keys import Keys  # 引入键盘对象，提供键盘按键支持（最后一个K要大写）
from selenium.webdriver.support.wait import WebDriverWait  # 浏览器等待对像
from selenium.webdriver.support import expected_conditions as EC  # 判断器
import urllib

# 1.导入了argparse这个包
import argparse
import io
# 2.包中的ArgumentParser类生成一个parser对象
#     a.ArgumentParser叫做参数解释器
#     b.其中的description描述这个参数解析器是干什么的
#     c.是当我们在命令行显示帮助信息的时候会看到description描述的信息。
parser = argparse.ArgumentParser(description="请输入图片的绝对路径！程序会依次执行，覆盖测试范围为迅捷图片前八个主要的功能。")
# 3.通过add_argument函数来增加参数
#     a.第一个"-p","-a","-w"是后面的简写
#     b.分别代表了path，amount和week
#     c.type是规范了这个参数的类型
#     d.default是默认值
parser.add_argument("-p","--path",type=str,default=None)
# 4. 最后采用对象的parse_args获取解析的参数
args = parser.parse_args()
print(args.path)

# 实例化

from selenium.webdriver import ActionChains  # 鼠标操作

import time
import os

# 获取Chrome驱动
browser = webdriver.Chrome()
# 文件路径
# base_dir = "/Users/mac/Desktop/"
# os.path.join(base_dir, "1.png")
path_to_image = str(args.path)
# 步骤计数器
pc = 0
def testBaiduSearch():
    '''
    模拟百度页面搜索
    :return:
    '''
    try:
        # 浏览器对象打开百度地址
        browser.get("https://www.baidu.com")
        # 查找id为 'kw'的标签，即输入框
        input = browser.find_element(by=By.ID, value="kw")
        # input = browser.find_element_by_id("kw")
        # 在输入框中填入'Python'
        input.send_keys("python")
        # '按下'回车键（第一种）
        # input.send_keys(Keys.ENTER)
        # 第二种，按按钮
        browser.find_element(by=By.ID, value="su").click()
        # browser.find_element_by_id("su").click()
        wait = WebDriverWait(browser, timeout=10, poll_frequency=0.5)
        wait.until(EC.presence_of_all_elements_located((By.ID, "content_left")))
    except Exception as e:
        print(e)
    else:
        # 打印请求的url
        print(browser.current_url)
        # 打印所有cookies
        print(browser.get_cookies())
    finally:
        # 等待10秒
        time.sleep(10)
        # 关闭浏览器对象
        browser.close()


def cutImage():
    '''
    模拟剪裁图片页面
    :return:
    '''
    try:
        # addImg=browser.find_element(by=By.XPATH,value="//div[@class='btn-first-load blue3 abs cp ct']")
        # 鼠标移动到某处单击
        # urllib.urlretrieve("http://www.peta.org/wp-content/uploads/2013/10/goat_2D00_list_2D00_1.jpg", path_to_image)
        # 前段时间在实现一个文件上传功能时发现谷歌等主流浏览器会出于安全考虑将上传文件的地址换为fakepath。比如c盘的某个文件，无论其原先处于哪个文件之下。在上传之后其会自动变为C\fakepath\文件名字。
        #
        #   解决方案如下：
        # upload_file = browser.find_element(by=By.CLASS_NAME, value="tui-image-editor-load-btn")
        # time.sleep(1)
        # upload_file.send_keys(path_to_image)
        # 上传图片
        input_image(path_to_image)
        # tooo = ActionChains(browser)
        # # tooo.send_keys("/Users/mac/Desktop/aaa2.png")
        # tooo.send_keys(Keys.ENTER)
        # ActionChains(browser).move_to_element(addImg).click(addImg).perform()

        # bu = browser.find_element(by=By.XPATH,value="div[@class='tui-image-editor-header-buttons']/button").click()
        # print(os.getcwd())  # 获得当前工作目录
        # upload_file= browser.find_element(by=By.CLASS_NAME,value='tui-image-editor-load-btn').send_keys("/Users/mac/Desktop/aaa2.png")
        # print(upload_file.get_attribute("value"))
        # pyk.type_string('/Users/mac/Desktop/aaa2.png')  # 打开文件所在目录，方便多个文件上传
        # time.sleep(1)
        # pyk.tap_key(pyk.return_key)
        # addImg.click()
        # time.sleep(1)
        # SendKeys
        # SendKeys.SendKeys('D:\\baidu.py')  # 发送文件地址
        # SendKeys.SendKeys("{ENTER}")  # 发送回车键

        # on = browser.find_element(by=By.CLASS_NAME,value="tui-image-editor").send_keys("/Users/mac/Desktop/aaa2.png")
        #
        # on.send_keys(Keys.ENTER)
        # 模拟选择区域
        canvas = browser.find_element(by=By.CLASS_NAME, value="lower-canvas")
        drawing = ActionChains(browser) \
            .click_and_hold(canvas) \
            .move_by_offset(100, 100) \
            .release()
        # # 执行链
        drawing.perform()

        # 剪裁图片
        browser.find_element(by=By.XPATH, value="//div[@class='fl cp ct btn-crop-apply blue3']").click()

        # 下载图片
        down_img()
        # 点击回车
        # print("--")
        # print("--")
        # ActionChains(browser).key_down(value=Keys.CONTROL,element=Keys.ENTER)
        # print("--")
    except Exception as e:
        print(e)
    else:
        # 打印请求的url
        print( browser.current_url)
        # 打印所有cookies
        print(browser.get_cookies())
    finally:
        # 隐形延时
        browser.implicitly_wait(8)


def addWord():
    try:
        # method1 = browser.find_element(by=By.XPATH, value="//div[@class='funlist white1']/div[1]/a[2]/li")
        # # 点击进入剪裁图片界面
        # method1.click()
        # 上传图片
        input_image(path_to_image)

        # 添加文字
        browser.find_element(by=By.XPATH, value="//img[@src='imgs/btn-wenzi-03.png']").click()
        # 在图片上点击
        time.sleep(1)
        canvas = browser.find_element(by=By.CLASS_NAME, value="lower-canvas")
        # 添加文字框
        addBox = ActionChains(browser) \
            .click(canvas)
        # # 执行链
        addBox.perform()
        time.sleep(1)
        # 双击
        word = ActionChains(browser).double_click() \
            .double_click() \
            .perform()
        word = ActionChains(browser).double_click() \
            .double_click() \
            .perform()
        time.sleep(1)
        # 添加文字
        key = ActionChains(browser)
        key.send_keys(Keys.NUMPAD6).perform()
        key.send_keys(Keys.NUMPAD6).perform()
        key.send_keys(Keys.NUMPAD6).perform()
        key.send_keys(Keys.ENTER).perform()
        time.sleep(1)
        # 变为红色字体
        browser.find_element(by=By.XPATH,
                             value="//ul[@class='color-picker-list color-picker-text white1']/li[1]").click()

        time.sleep(1)
        # 改变大小
        rr = browser.find_element(by=By.XPATH, value="//input[@class='range-text-size cp abs']")
        rr.click()
        time.sleep(1)
        # 撤销
        browser.find_element(by=By.XPATH, value="//img[@data='chexiao']").click()
        time.sleep(1)
        # 删除
        browser.find_element(by=By.XPATH, value="//img[@data='delete']").click()

    except Exception as e:
        print(e)
    else:
        print(browser.current_url)
        # 打印所有cookies
        print(browser.get_cookies())
    finally:
        # 隐形延时
        browser.implicitly_wait(8)


def addShape():
    try:
        # global pc
        # pc = pc +1
        # print(str(pc)+"进入添加图形界面")
        # method1 = browser.find_element(by=By.XPATH, value="//div[@class='funlist white1']/div[1]/a[3]/li")
        # # 点击进入剪裁图片界面
        # method1.click()
        # 上传图片
        input_image(path_to_image)

        # 选择实心矩形
        browser.find_element(by=By.XPATH,
                             value="//li[@data='rect1']/p").click()
        # 选择红
        browser.find_element(by=By.XPATH,
                             value="//ul[@class='color-picker-list color-picker-shape white1']/li[1]").click()
        time.sleep(1)

        canvas = browser.find_element(by=By.CLASS_NAME,
                                      value="lower-canvas")
        # 拖动

        fullRect = ActionChains(browser) \
            .click_and_hold(canvas) \
            .move_by_offset(100, 100) \
            .release()
        fullRect.perform()
        # 选择白黄色
        browser.find_element(by=By.XPATH,
                             value="//ul[@class='color-picker-list color-picker-shape white1']/li[3]").click()
        time.sleep(1)
        # 移动矩形
        move = ActionChains(browser) \
            .click_and_hold(canvas) \
            .move_by_offset(-100, 100) \
            .move_by_offset(140, -150) \
            .release()
        move.perform()
        time.sleep(1)
        # 选择空心矩形
        browser.find_element(by=By.XPATH,
                             value="//li[@data='rect2']/p").click()

        # 选择绿色
        browser.find_element(by=By.XPATH,
                             value="//ul[@class='color-picker-list color-picker-shape white1']/li[4]").click()
        emptyRect = ActionChains(browser) \
            .click_and_hold(canvas) \
            .move_by_offset(-100, 100) \
            .move_by_offset(140, -150) \
            .release()
        emptyRect.perform()
        time.sleep(1)
        # 选择实心圆形
        browser.find_element(by=By.XPATH,
                             value="//li[@data='circle1']/p").click()

        # 选择蓝色
        browser.find_element(by=By.XPATH,
                             value="//ul[@class='color-picker-list color-picker-shape white1']/li[5]").click()
        fullCircle = ActionChains(browser) \
            .click_and_hold(canvas) \
            .move_by_offset(-150, -100) \
            .move_by_offset(100, 130) \
            .release()
        fullCircle.perform()
        time.sleep(1)
        # 选择空心圆形
        browser.find_element(by=By.XPATH,
                             value="//li[@data='circle2']/p").click()

        # 选择紫色
        browser.find_element(by=By.XPATH,
                             value="//ul[@class='color-picker-list color-picker-shape white1']/li[2]").click()
        emptyCircle = ActionChains(browser) \
            .click_and_hold(canvas) \
            .move_by_offset(-100, -100) \
            .move_by_offset(-100, -100) \
            .release()
        emptyCircle.perform()
        time.sleep(1)
        # 撤销
        browser.find_element(by=By.XPATH,
                             value="//img[@data='chexiao']").click()
        time.sleep(1)
        # 删除
        browser.find_element(by=By.XPATH,
                             value="//img[@data='delete']").click()


    except Exception as e:
        print(e)
    else:
        print(browser.current_url)
        # 打印所有cookies
        print(browser.get_cookies())
    finally:
        # 隐形延时
        browser.implicitly_wait(8)


def draw():
    try:
        # method1 = browser.find_element(by=By.XPATH,
        #                                value="//div[@class='funlist white1']/div[1]/a[4]/li")
        # # 点击进入剪裁图片界面
        # method1.click()
        # 上传图片
        input_image(path_to_image)
        # 依次绘制不同颜色，不同粗细的线条
        # 点击红色
        browser.find_element(by=By.XPATH,
                             value="//ul[@class='color-picker-list color-picker-draw white1']/li[1]").click()

        canvas=browser.find_element(by=By.CLASS_NAME,
                                    value="lower-canvas")
        drawPen1 = ActionChains(browser) \
            .click(canvas) \
            .drag_and_drop_by_offset(canvas, -100, -100) \
            .click_and_hold(canvas) \
            .move_by_offset(100, -100) \
            .release()
        drawPen1.perform()
        time.sleep(1)
        # 点击黄色
        browser.find_element(by=By.XPATH,
                             value="//ul[@class='color-picker-list color-picker-draw white1']/li[3]").click()
        # 改变宽度
        browser.find_element(by=By.XPATH,
                             value="//div[@class='range-rect white1']/input[@class='range-brush-width cp abs']").click()
        drawPen2 = ActionChains(browser) \
            .click(canvas) \
            .drag_and_drop_by_offset(canvas, -100, -50) \
            .click_and_hold(canvas) \
            .move_by_offset(100, -50) \
            .release()
        drawPen2.perform()
        time.sleep(1)
        # 点击绿色
        browser.find_element(by=By.XPATH,
                             value="//ul[@class='color-picker-list color-picker-draw white1']/li[4]").click()
        # 改变宽度
        width= browser.find_element(by=By.XPATH,
                                    value="//div[@class='range-rect white1']/input[@class='range-brush-width cp abs']")
        width.click()
        width.send_keys(150)
        drawPen3 = ActionChains(browser) \
            .click(canvas) \
            .drag_and_drop_by_offset(canvas, -100, 0) \
            .click_and_hold(canvas) \
            .move_by_offset(100, 0) \
            .release()
        drawPen3.perform()
        time.sleep(1)
        # 撤销
        browser.find_element(by=By.XPATH,value="//img[@data='chexiao']").click()

        # 重做
        browser.find_element(by=By.XPATH, value="//img[@data='fuwei']").click()
        time.sleep(1)

    except Exception as e:
        print(e)
    else:
        print(browser.current_url)
        # 打印所有cookies
        print(browser.get_cookies())

    finally:
        # 隐性延时
        browser.implicitly_wait(8)


def addFilter():
    try:

        # # 确保是进入的本功能页面
        # method1 = browser.find_element(by=By.XPATH,
        #                                value="//div[@class='funlist white1']/div[3]/a[1]/li")
        # method1.click()
        # 上传图片
        input_image(path_to_image)
        # 循环换滤镜
        for num in range(2,9):
            browser.find_element(by=By.XPATH,value="//section[@class='fl menu-filter']/li["+str(num)+"]").click()
            time.sleep(1)



    except Exception as e:
        print(e)
    else:
        print(browser.current_url)
        # 打印所有cookies
        print(browser.get_cookies())
    finally:
        browser.implicitly_wait(8)

def rotateImage():
    try:

        # # 确保是进入的本功能页面
        # method1 = browser.find_element(by=By.XPATH,
        #                                value="//div[@class='funlist white1']/div[3]/a[2]/li")
        # method1.click()
        # 上传图片
        input_image(path_to_image)
        # 左旋转
        time.sleep(1)
        browser.find_element(by=By.XPATH,value="//li[@class='fl ct btn-rotate-ccw cp']/p").click()
        # 右旋转
        time.sleep(1)
        browser.find_element(by=By.XPATH,value="//li[@class='fl ct btn-rotate-cw cp']/img[@src='imgs/btn-shun-02.png']").click()
        time.sleep(1)
        # 无极旋转
        rotate= browser.find_element(by=By.XPATH,value="//div[@class='range-rect2 white1']/input[@class='range-rotate-full cp abs']")
        move_rotate = ActionChains(browser)\
            .click_and_hold(rotate)\
            .move_by_offset(90,0)\
            .release()
        move_rotate.perform()
        time.sleep(1)
    except Exception as e:
        print(e)
    else:
        print(browser.current_url)
        # 打印所有cookies
        print(browser.get_cookies())
    finally:
        browser.implicitly_wait(8)

def imageMirror():
    try:
        # # 确保是进入的本功能页面
        # method1 = browser.find_element(by=By.XPATH,
        #                                value="//div[@class='funlist white1']/div[3]/a[3]/li")
        # method1.click()
        # 上传图片
        input_image(path_to_image)
        time.sleep(1)
        # 垂直镜像
        browser.find_element(by=By.XPATH,value="//li[@class='fl ct btn-flip-y']/p").click()
        time.sleep(1)
        # 水平镜像
        browser.find_element(by=By.XPATH, value="//li[@class='fl ct btn-flip-x']/p").click()
        # 下载图片
        time.sleep(1)
    except Exception as e:
        print(e)
    else:
        print(browser.current_url)
        # 打印所有cookies
        print(browser.get_cookies())
    finally:
        browser.implicitly_wait(8)

def editImage():
    try:
        # # 确保是进入的本功能页面
        # method1 = browser.find_element(by=By.XPATH,
        #                                value="//div[@class='funlist white1']/div[3]/a[4]/li")
        # method1.click()
        # 上传图片
        input_image(path_to_image)
        # 涂鸦
        browser.find_element(by=By.XPATH,value="//section[@class='tool-bar-func abs']/li[1]").click()
        draw()
        # 添加形状
        time.sleep(1)
        browser.find_element(by=By.XPATH,value="//section[@class='tool-bar-func abs']/li[2]").click()
        addShape()
        time.sleep(1)
        # 添加文字
        browser.find_element(by=By.XPATH,value="//section[@class='tool-bar-func abs']/li[3]").click()
        addWord()
        time.sleep(1)
        # 剪裁图片
        browser.find_element(by=By.XPATH,value="//section[@class='tool-bar-func abs']/li[5]").click()
        cutImage()
        time.sleep(1)
        # 旋转图片
        browser.find_element(by=By.XPATH,value="//section[@class='tool-bar-func abs']/li[6]").click()
        rotateImage()
        time.sleep(1)
        # 图片镜像
        browser.find_element(by=By.XPATH,value="//section[@class='tool-bar-func abs']/li[7]").click()
        imageMirror()

    except Exception as e:
        print(e)
    else:
        print(browser.current_url)
        # 打印所有cookies
        print(browser.get_cookies())
    finally:
        browser.implicitly_wait(8)

def onlinePS():
    pass


def imageEffects():
    pass

def input_image(path):
    '''
    封装为一个方法，上传图片
    :param path:
    :return:
    '''
    input = browser.find_element(by=By.CLASS_NAME, value="tui-image-editor-load-btn")
    time.sleep(1)
    input.send_keys(path)

def save_and_quit():
    '''
    封装为一个方法，保存退出
    :return:
    '''
    # 实现下个功能
    browser.find_element(by=By.XPATH, value="//p[@class='fl cp txtBlue']/../following-sibling::li[1]/p").click()
    # 确认
    browser.find_element(by=By.XPATH, value="//div[@class='fr btn-hint-go blue3 cp ct']").click()

def down_img():
    '''
    封装方法，下载图片
    :return:
    '''
    # 下载图片
    browser.find_element(by=By.XPATH, value="//div[@class='fl cp ct btn-down-load']").click()
    time.sleep(1)

def main():
    browser.get("http://app.xunjietupian.com/")
    # 应用自顶向下的设计思路，首先写出大概的程序轮廓，然后对每个程序模块一一实现
    # 因为a被li所占用，所以直接点击li即可跳转
    method1 = browser.find_element(by=By.XPATH, value="//div[@class='funlist white1']/div[1]/a[1]/li")
    # # 点击进入剪裁图片界面
    method1.click()
    # 定位到添加图片的div
    cutImage()
    # 下载图片
    down_img()
    # 保存退出
    save_and_quit()
    time.sleep(1)
    addWord()
    # 下载图片
    down_img()
    save_and_quit()
    time.sleep(1)
    addShape()
    # 下载图片
    down_img()
    save_and_quit()
    time.sleep(1)
    draw()
    # 下载图片
    down_img()
    save_and_quit()
    time.sleep(1)
    addFilter()
    # 下载图片
    down_img()
    save_and_quit()
    time.sleep(1)
    rotateImage()
    down_img()
    save_and_quit()
    time.sleep(1)
    # 下载图片
    imageMirror()
    # 下载图片
    down_img()
    save_and_quit()
    editImage()
    # browser.find_element(by=By.XPATH, value="")
    # onlinePS()
    # browser.find_element(by=By.XPATH, value="")
    # imageEffects()

    time.sleep(1)
    browser.close()


if __name__ == '__main__':
    main()
