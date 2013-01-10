GitHub Doc

如果你喜欢Go，如果你想用Go解决http server的问题

请拜读谢大的《Go Web 编程》https://github.com/astaxie/build-web-application-with-golang

------------------

如果你像我一样，只用iPad看书读代码，但又是Wifi Only

那么这个拙劣的工具可以把.md和相关图片下载为.html

好吧，其实我就是想离线读《Go Web 编程》，而且这个系列极具价值，可以作为参考手册

------------------

go install githubdoc

githubdoc.exe https://github.com/astaxie/build-web-application-with-golang

------------------

更新：

2013-1-10
    加入goroutine（用Go怎么能不用routine），效率得到较大提升

    大家可以在 parser.go中对routines pool的大小进行设置，影响还是挺大的，在本机的环境下测试
    routine count | time
        4            41s
        8            27s
       16            17s

    但没有测试在 GOMAXPROCS下的效果

------------------

下一步：<br>
    增量更新<br>
    研究GitHub的Css，争取离线的.html达到原文的效果<br>
    研究Go导出PDF<br>
