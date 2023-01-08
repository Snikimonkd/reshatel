set terminal png
set grid
set xrange [-10:20]
set yrange [-10:20]
set output 'out.png'
plot 'data' using 1:2:($3-$1):($4-$2) with vectors nohead