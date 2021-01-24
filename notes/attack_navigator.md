# ATT&CK® Navigator Learning Notes

## 层文件格式

- MITRE ATT&CK Navigator 层文件存储格式为 JSON
- 以 JSON 格式从 [ATT&CK® Navigator](https://mitre-attack.github.io/attack-navigator/) 中导出任意编辑后的矩阵，由于技能点不能增删，JSON 文件将会标记使用的矩阵模板，并记录编辑后与原矩阵的差异部分
    ```js
    {
        "name": "layer",
        "versions": {
            "attack": "8",
            "navigator": "4.0",
            "layer": "4.0"
        },
        "domain": "enterprise-attack",  // 矩阵代表的科技领域
        "description": "",
        "filters": {
            "platforms": [  // 指定平台，只展示有这些平台标签的技术
                "Linux",
                "macOS",
                "Windows",
                "Office 365",
                "Azure AD",
                "AWS",
                "GCP",
                "Azure",
                "SaaS",
                "PRE",
                "Network"
            ]
        },
        "sorting": 0,
        "layout": {
            "layout": "side",
            "showID": false,
            "showName": true
        },
        "hideDisabled": false,
        "techniques": [
            {
                "techniqueID": "T1480",
                "tactic": "defense-evasion",
                "color": "#e60d0d",
                "comment": "",
                "enabled": true,
                "metadata": [],
                "showSubtechniques": false
            }
        ],
        "gradient": {
            "colors": [ // 标记分值的渐变色
                "#ff6666",
                "#ffe766",
                "#8ec843"
            ],
            "minValue": 0,  // 分数最小值
            "maxValue": 100 // 分数最大值
        },
        "legendItems": [],
        "metadata": [],
        "showTacticRowBackground": false,
        "tacticRowBackground": "#dddddd",
        "selectTechniquesAcrossTactics": true,
        "selectSubtechniquesWithParent": false
    }
    ```
- 若制作能增删技能点的矩阵，仍然以特定的科技领域为基础，但需要增加额外的标记

### 参考资料

[ATT&CK™ Navigator Layer File Format Definition](https://mitre-attack.github.io/attack-navigator/assets/NavigatorLayerFileFormatv4.pdf)

## Simple Layers

### 利用 CSV 文件生成层文件（JSON）

- 提供的样例：[simple_input.csv](https://github.com/mitre-attack/attack-navigator/blob/master/layers/data/csv/simple_input.csv)
  - 包含列：`TechID`、`Software`、`Groups`、`References`
- 脚本 [attack_layers_simple.py](https://github.com/mitre-attack/attack-navigator/blob/master/layers/attack_layers/attack_layers_simple.py) 比较简单，读入 CSV 文件，计算分值（使用数值、计算的方法自定义）并添加到记录层信息的变量中


## Mitre Att&ck Tips
- As long as a sub-techniques conceptually falls under a technique
(e.g. sub-techniques that are conceptually a type of process injection will be under process
injection), each sub-technique can contribute to which tactics a technique is a part of but are not
required to fulfill every parent technique’s tactic (i.e. the Process Hollowing sub-technique can
be used for Defense Evasion but not Privilege Escalation even though the Process Injection
technique covers both tactics).