#!/bin/sh

agent="wget --max-redirect 0 -q -O -"

if [ "$1" == "-h" ] || [ "$1" == "--help" ]; then
	cat<<EOF
Usage: $0 [BEGIN] [END]

Fetch Project Euler problem statements from BEGIN to END included.
Save them to

	<n>/<n>.html

where <n> is the number of the problem.

With no arguments, fetch them all.
If END is unspecified, fetch only BEGIN.
If END < BEGIN, fetch from BEGIN to the last one.

If file exists, statement is not fetched.

EOF
	exit
fi

root="$(realpath "$0")"
root="${root%/*}"

pb=1
pb_end=0
[ -n "$1" ] && { pb="$1"; pb_end=$pb; }
[ -n "$2" ] && pb_end="$2"

if [ $pb_end -lt $pb ]; then
	## If pb_end <= pb-1, we want to fetch everything.
	pb_end=0
else
	## If pb_end >= pb, we want to fetch up to pb_end included.
	pb_end=$(($pb_end+1))
fi

while [ $pb -ne $pb_end ] ; do
	target=$pb
	pb=$(($pb+1))
	[ -f "$root/$target/$target.html" ] && continue
	buf="$($agent "https://projecteuler.net/problem=$target")"
	[ $? -ne 0 ] && break

	[ ! -d "$root/$target" ] && mkdir "$root/$target"
	[ ! -d "$root/$target" ] && continue
	echo "$root/$target/$target.html"
	echo "$buf" | awk '/<h2>.*<\/h2>/ {sub(/<\/h2>.*/, "</h2>"); print}; /<div class="problem_content" role="problem">/ {jump=1;next}; jump == 1 && $0 !~ "</div><br />" {gsub("<a href=\"project/resources/", "<a href=\""); print; next}; jump == 1 {exit}' > "$root/$target/$target.html"

	## Save attached files.
	while IFS= read -r i; do
		[ -n "$i" ] &&	$agent "https://projecteuler.net/project/resources/$i" > "$root/$target/$i"
		done <<EOF
$(echo "$buf" | awk -F '<a href="project/resources/' 'NF > 1 {for (i=2; i<=NF; i++) {gsub(/">.*/, "", $i); print $i}}')
EOF
done
