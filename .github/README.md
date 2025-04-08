# PixCluster - 像素聚类智能分析平台

本项目为数据可视化课程大作业，主要功能为对图片的像素值进行聚类并将结果使用数据图表进行可视化

<p align="center">
  <img src="https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=fff&style=flat-square" alt="Go Badge">
  <img src="https://img.shields.io/badge/TypeScript-3178C6?logo=typescript&logoColor=fff&style=flat-square" alt="TypeScript Badge">
  <img src="https://img.shields.io/badge/Make-6D00CC?logo=make&logoColor=fff&style=flat-square" alt="Make Badge">
  <img src="https://img.shields.io/badge/React-61DAFB?logo=react&logoColor=000&style=flat-square" alt="React Badge">
  <img src="https://img.shields.io/badge/Next.js-000?logo=nextdotjs&logoColor=fff&style=flat-square" alt="Next.js Badge">
  <img src="https://img.shields.io/badge/Ant%20Design-0170FE?logo=antdesign&logoColor=fff&style=flat-square" alt="Ant Design Badge">
  <img src="https://img.shields.io/badge/AntV-8B5DFF?logo=antv&logoColor=fff&style=flat-square" alt="AntV Badge">
  <img src="https://img.shields.io/badge/Alibaba%20Cloud-FF6A00?logo=alibabacloud&logoColor=fff&style=flat-square" alt="Alibaba Cloud Badge">
</p>

## 🌐 网站截图

<p align="center"><img src="../images/website.png" alt="Website screenshot" width="2304"></p>

## ✨ 特色功能

### 🚀 高效聚类后端引擎
- 基于 **Go语言** 构建高性能后端接口
- 依托于 **函数计算** 基础设施构建高可用、高弹性、低成本的服务
- 采用 **KMeans++ 聚类算法** 实现图片像素级颜色聚类分析
- 通过 **Goroutine 协程并发** 优化计算效率，支持高并发图片处理任务
- 支持多种格式的图片解析，自动提取RGB/HEX色彩值进行聚类分析

### 🌀 AI文生图与智能分析
- 集成 **通义万相API** 实现文生图功能，用户输入prompt即可生成创意图片
- 生成图片自动触发聚类分析流水线，实时返回颜色分布数据
- 调用 **通义千问视觉模型** 对聚类结果进行多维度解读，生成图文分析报告

### 🎨 交互式数据可视化
- 使用 **Ant Design Charts** 构建动态柱状图，直观展示颜色簇分布比例
- 支持图表颜色映射预览，鼠标悬浮可高亮对应颜色簇
- 提供聚类中心具体HEX值，方便设计师直接取用配色方案

### 🖥️ 全场景响应式布局
- 采用响应式布局方案，完美适配桌面/平板/手机等设备
- 基于CSS断点动态重组界面，小屏幕自动省略次要内容
- 图片上传区智能缩放，生成结果自适应屏幕分辨率

### 🧬 智能化工作流
- 从「文字→生成图片→聚类分析→可视化→智能报告」全流程自动化
- 各模块间通过异步函数进行执行，支持分析结果逐步展示
- 错误重试机制与加载状态优化，保障长流程操作体验

## ⚡ 快速开始

### 克隆仓库

```bash
git clone https://github.com/fenggwsx/PixCluster.git && cd PixCluster
```

### 前端开发

#### 安装项目依赖

```bash
pnpm --prefix web install
pnpm --prefix proxy install
```

#### 启动开发服务器

```bash
pnpm --prefix web dev
```

#### 配置代理服务器

> [!NOTE]
> 代理服务器用于解决开发环境下的跨域问题。

1. 复制环境模板文件：

```bash
cp proxy/.env.example proxy/.env.local
```

2. 编辑`proxy/.env.local`文件，配置实际的前后端地址：

```env
# Frontend
FRONTEND_URL=http://localhost:3000

# Backend
BACKEND_URL=https://example.com
```

> [!NOTE]
> 后端暂不支持本地调试，需将后端服务部署后在前端中进行调试。

3. 启动代理服务器：
```bash
pnpm --prefix proxy start
```

> [!TIP]
> 若需经常使用代理服务器，可以将代理服务器置于后台进程中。

#### 访问应用

打开浏览器访问：[http://localhost:9000](http://localhost:9000)

### 项目构建

本项目基于阿里云函数计算，需将代码编译并部署至函数计算平台。

> [!NOTE]
> 函数计算（Function Compute）是一个事件驱动的全托管 Serverless 计算服务。

#### 服务概览

| 服务名称     | 功能描述     | 入口函数代码文件路径     |
| ------------ | ------------ | ------------------------ |
| `web`        | 前端界面     | `web/app/layout.tsx`     |
| `text2image` | 文生图服务   | `cmd/text2image/main.go` |
| `kmeans`     | 像素聚类分析 | `cmd/kmeans/main.go`     |
| `summarize`  | 文本智能总结 | `cmd/summarize/main.go`  |

#### 编译单个服务

执行以下命令编译指定服务（如`kmeans`）：

```bash
make build-kmeans
```

- 输出文件：`build/kmeans.zip`
- 其他服务替换命令中的服务名，如`build-text2image`

#### 编译所有后端服务

执行以下命令编译所有后端服务：

```bash
make build-apps
```

- 输出文件：`build`目录下生成`text2image.zip`, `kmeans.zip`, `summarize.zip`

#### 编译所有服务

执行以下命令编译所有服务：

```bash
make
```

- 输出文件：`build`目录下生成`web.zip`, `text2image.zip`, `kmeans.zip`, `summarize.zip`

### 项目部署

在阿里云函数计算平台新建函数，进行相应的配置，最后上传生成的`zip`压缩包即可完成部署。

#### 前端配置

对于前端服务，需创建Web函数，并添加自定义层Nginx，启动命令如下：

```bash
/opt/bin/nginx -c /code/nginx.conf -g 'daemon off;'
```

#### 后端配置

对于后端服务，需创建事件函数，运行时为Golang，并分配合适的CPU和内存。
