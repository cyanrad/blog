(ls -s static\markdown\).name | each {|f| 
    let i = 'static\markdown\' + $f 
    let o = 'static\posts\' + ($f | str replace '.md' '.html')
    pandoc -i $i -o $o
}
