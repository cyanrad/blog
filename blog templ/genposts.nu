(ls -s content\markdown\).name | each {|f| 
    let i = 'content\markdown\' + $f 
    let o = 'content\posts\' + ($f | str replace '.md' '.html')
    pandoc -i $i -o $o
}
