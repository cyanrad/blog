(ls posts).name | each {|f| pandoc -o $"static\\($f).html" $f}
