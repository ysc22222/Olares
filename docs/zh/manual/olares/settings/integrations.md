---
outline: [2, 3]
description: 将 Olares Space 与第三方服务连接，扩展系统功能。了解如何在 Settings 中查看、授权及管理集成服务，实现数据的无缝同步。
---

# 在设置中管理集成

设置中的**集成**分区集中展示所有已连接到 Olares 的第三方服务，并支持使用 API 凭证手动配置云对象存储。

:::tip 注意
OAuth 类型的集成以及 Olares Space 需在**LarePass**应用中完成连接，详见 [LarePass 集成文档](../../larepass/integrations.md)。
:::
## 查看与管理现有集成

1. 通过 Dock 或启动台打开**设置**。  
2. 在左侧菜单选择**集成**，即可看到已授权的服务列表。  
3. 点击任一集成卡片查看连接状态和操作选项。  
4. 在**账户设置**页面点击**删除**可移除该集成。  

## 通过 API 密钥添加云对象存储

Olares 支持手动配置 **AWS S3** 和**腾讯云 COS**：

1. 进入**设置** > **集成**，点击右上角**添加账户**。  
2. 选择 **AWS S3** 或 **Tencent COS**，点击**确认**。  
3. 在弹出的挂载对话框输入以下信息：  
   - **Access Key**  
   - **Secret Key**  
   - **Region**  
   - **Bucket name**  
4. 点击**下一步**。凭证验证通过后将显示成功提示。

你也可以在 [LarePass](../../larepass/integrations.md#通过-api-密钥添加云盘) 中完成同样的配置。
