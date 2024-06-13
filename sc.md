[SC1000] `$` is not used specially and should therefore be escaped.
[SC1001] This `\o` will be a regular 'o' in this context.
[SC1003] Want to escape a single quote? `echo 'This is how it'\''s done'`.
[SC1004] This backslash+linefeed is literal. Break outside single quotes if you just want to break the line.
[SC1007] Remove space after `=` if trying to assign a value (or for empty string, use `var=''` ... ).
[SC1008] This shebang was unrecognized. ShellCheck only supports sh/bash/dash/ksh. Add a 'shell' directive to specify.
[SC1009] The mentioned parser error was in ...
[SC1010] Use semicolon or linefeed before `done` (or quote to make it literal).
[SC1011] This apostrophe terminated the single quoted string!
[SC1012] `\t` is just literal `t` here. For tab, use `"$(printf '\t')"` instead.
[SC1014] Use `if cmd; then ..` to check exit code, or `if [ "$(cmd)" = .. ]` to check output.
[SC1015] This is a unicode double quote. Delete and retype it.
[SC1016] This is a Unicode single quote. Delete and retype it.
[SC1017] Literal carriage return. Run script through `tr -d '\r'` .
[SC1018] This is a unicode non-breaking space. Delete it and retype as space.
[SC1019] Expected this to be an argument to the unary condition.
[SC1020] You need a space before the `]` or `]]`
[SC1026] If grouping expressions inside `[[..]]`, use `( .. )`.
[SC1027] Expected another argument for this operator.
[SC1028] In `[..]` you have to escape `\( \)` or preferably combine `[..]` expressions.
[SC1029] In `[[..]]` you shouldn't escape `(` or `)`.
[SC1033] Test expression was opened with double `[[` but closed with single `]`. Make sure they match.
[SC1034] Test expression was opened with double `[` but closed with single `]]`. Make sure they match.
[SC1035] You need a space here
[SC1036] `(` is invalid here. Did you forget to escape it?
[SC1037] Braces are required for positionals over 9, e.g. `${10}`.
[SC1038] Shells are space sensitive. Use `< <(cmd)`, not `<<(cmd)`.
[SC1039] Remove indentation before end token (or use `<<-` and indent with tabs).
[SC1040] When using `<<-`, you can only indent with tabs.
[SC1041] Found `eof` further down, but not on a separate line.
[SC1043] Found EOF further down, but with wrong casing.
[SC1044] Couldn't find end token `EOF` in the here document.
[SC1045] It's not `foo &; bar`, just `foo & bar`.
[SC1046] Couldn't find `fi` for this `if`.
[SC1047] Expected `fi` matching previously mentioned `if`.
[SC1048] Can't have empty then clauses (use `true` as a no-op).
[SC1049] Did you forget the `then` for this `if`?
[SC1050] Expected `then`.
[SC1051] Semicolons directly after `then` are not allowed. Just remove it.
[SC1052] Semicolons directly after `then` are not allowed. Just remove it.
[SC1053] Semicolons directly after `else` are not allowed. Just remove it.
[SC1054] You need a space after the `{`.
[SC1055] You need at least one command here. Use `true;` as a no-op.
[SC1056] Expected a `}`. If you have one, try a `;` or `\n` in front of it.
[SC1057] Did you forget the `do` for this loop?
[SC1058] Expected `do`.
[SC1059] Semicolon is not allowed directly after `do`. You can just delete it.
[SC1060] Can't have empty do clauses (use `true` as a no-op)
[SC1061] Couldn't find `done` for this `do`.
[SC1062] Expected `done` matching previously mentioned `do`.
[SC1063] You need a line feed or semicolon before the `do`.
[SC1064] Expected a `{` to open the function definition.
[SC1065] Trying to declare parameters? Don't. Use `()` and refer to params as `$1`, `$2`, ..
[SC1066] Don't use `$` on the left side of assignments.
[SC1067] For indirection, use arrays, `declare "var$n=value"`, or (for sh) `read`/`eval`
[SC1068] Don't put spaces around the `=` in assignments.
[SC1069] You need a space before the `[`.
[SC1070] Parsing stopped here. Mismatched keywords or invalid parentheses?
[SC1071] ShellCheck only supports sh/bash/dash/ksh scripts. Sorry!
[SC1072] Unexpected ..
[SC1073] Couldn't parse this (thing). Fix to allow more checks.
[SC1074] Did you forget the `;;` after the previous case item?
[SC1075] Use `elif` instead of `else if`.
[SC1076] Trying to do math? Use e.g. `[ $((i/2+7)) -ge 18 ]`.
[SC1077] For command expansion, the tick should slant left (`` ` `` vs `´`).
[SC1078] Did you forget to close this double quoted string?
[SC1079] This is actually an end quote, but due to next char it looks suspect.
[SC1080] You need `\` before line feeds to break lines in `[ ]`.
[SC1081] Scripts are case-sensitive. Use `if`, not `If`.
[SC1082] This file has a UTF-8 BOM. Remove it with: `LC_CTYPE=C sed '1s/^...//' < yourscript`.
[SC1083] This `{`/`}` is literal. Check if `;` is missing or quote the expression.
[SC1084] Use `#!`, not `!#`, for the shebang.
[SC1086] Don't use `$` on the iterator name in for loops.
[SC1087] Use braces when expanding arrays, e.g. `${array[idx]}` (or `${var}[..` to quiet).
[SC1088] Parsing stopped here. Invalid use of parentheses?
[SC1089] Parsing stopped here. Is this keyword correctly matched up?
[SC1090] Can't follow non-constant source. Use a directive to specify location
[SC1091] Not following: (error message here)
[SC1092] Stopping at 100 `source` frames :O
[SC1094] Parsing of sourced file failed. Ignoring it.
[SC1095] You need a space or linefeed between the function name and body.
[SC1097] Unexpected `==`. For assignment, use `=`. For comparison, use `[`/`[[`.
[SC1098] Quote/escape special characters when using `eval`, e.g. `eval "a=(b)"`.
[SC1099] You need a space before the `#`.
[SC1100] This is a unicode dash. Delete and retype as ASCII minus.
[SC1101] Delete trailing spaces after `\` to break line (or use quotes for literal space).
[SC1102] Shells disambiguate `$((` differently or not at all. For `$(command substitution)`, add space after `$(` . For `$((arithmetics))`, fix parsing errors.
[SC1103] This shell type is unknown. Use e.g. `sh` or `bash`.
[SC1104] Use `#!`, not just `!`, for the shebang.
[SC1105] Shells disambiguate `((` differently or not at all. If the first `(` should start a subshell, add a space after it.
[SC1106] In arithmetic contexts, use `<` instead of `-lt`
[SC1107] This directive is unknown. It will be ignored.
[SC1108] You need a space before and after the `=` .
[SC1109] This is an unquoted HTML entity. Replace with corresponding character.
[SC1110] This is a unicode quote. Delete and retype it (or quote to make literal).
[SC1111] This is a unicode quote. Delete and retype it (or ignore/singlequote for literal).
[SC1112] This is a unicode quote. Delete and retype it (or ignore/doublequote for literal).
[SC1113] Use `#!`, not just `#`, for the shebang.
[SC1114] Remove leading spaces before the shebang.
[SC1115] Remove spaces between `#` and `!` in the shebang.
[SC1116] Missing `$` on a `$((..))` expression? (or use `( (` for arrays).
[SC1117] Backslash is literal in `"\n"`. Prefer explicit escaping: `"\\n"`.
[SC1118] Delete whitespace after the here-doc end token.
[SC1119] Add a linefeed between end token and terminating `)`.
[SC1120] No comments allowed after here-doc token. Comment the next line instead.
[SC1121] Add `;`/`&` terminators (and other syntax) on the line with the `<<`, not here.
[SC1122] Nothing allowed after end token. To continue a command, put it on the line with the `<<`.
[SC1123] ShellCheck directives are only valid in front of complete compound commands, like `if`, not e.g. individual `elif` branches.
[SC1124] ShellCheck directives are only valid in front of complete commands like `case` statements, not individual case branches.
[SC1125] Invalid `key=value` pair in directive
[SC1126] Place shellcheck directives before commands, not after.
[SC1127] Was this intended as a comment? Use `#` in sh.
[SC1128] The shebang must be on the first line. Delete blanks and move comments.
[SC1129] You need a space before the `!`.
[SC1130] You need a space before the :.
[SC1131] Use `elif` to start another branch.
[SC1132] This `&` terminates the command. Escape it or add space after `&` to silence.
[SC1133] Unexpected start of line. If breaking lines, `|`/`||`/`&&` should be at the end of the previous one.
[SC1134] Error parsing `shellcheckrc`:
[SC1135] Prefer escape over ending quote to make `$` literal. Instead of `"It costs $"5`, use `"It costs \$5"`
[SC1136] Unexpected characters after terminating `]`. Missing semicolon/linefeed?
[SC1137] Missing second `(` to start arithmetic for ((;;)) loop
[SC1138] Shells are space sensitive. Use `< <(cmd)`, not `<< (cmd)`.
[SC1139] Use `||` instead of `-o` between test commands.
[SC1140] Unexpected parameters after condition. Missing `&&`/`||`, or bad expression?
[SC1141] Unexpected tokens after compound command. Bad redirection or missing `;`/`&&`/`||`/`|`?
[SC1142] Use `done < <(cmd)` to redirect from process substitution (currently missing one `<`).
[SC1143] This backslash is part of a comment and does not continue the line.
[SC1144] `external-sources` can only be enabled in .shellcheckrc, not in individual files.
[SC1145] Unknown `external-sources` value. Expected `true`/`false`.
[SC2000] See if you can use `${#variable}` instead
[SC2001] See if you can use `${variable//search/replace}` instead.
[SC2002] Useless cat. Consider `cmd < file | ..` or `cmd file | ..` instead.
[SC2003] expr is antiquated. Consider rewriting this using `$((..))`, `${}` or `[[  ]]`.
[SC2004] `$`/`${}` is unnecessary on arithmetic variables.
[SC2005] Useless `echo`? Instead of `echo $(cmd)`, just use `cmd`
[SC2006] Use `$(...)` notation instead of legacy backticked `` `...` ``.
[SC2007] Use `$((..))` instead of deprecated `$[..]`.
[SC2008] `echo` doesn't read from stdin, are you sure you should be piping to it?
[SC2009] Consider using `pgrep` instead of grepping `ps` output.
[SC2010] Don't use `ls | grep`. Use a glob or a for loop with a condition to allow non-alphanumeric filenames.
[SC2011] Use `find -print0` or `find -exec` to better handle non-alphanumeric filenames.
[SC2012] Use `find` instead of `ls` to better handle non-alphanumeric filenames.
[SC2013] To read lines rather than words, pipe/redirect to a `while read` loop.
[SC2014] This will expand once before find runs, not per file found.
[SC2015] Note that `A && B || C` is not if-then-else. C may run when A is true.
[SC2016] Expressions don't expand in single quotes, use double quotes for that.
[SC2017] Increase precision by replacing `a/b*c` with `a*c/b`.
[SC2018] Use `[:lower:]` to support accents and foreign alphabets.
[SC2019] Use `[:upper:]` to support accents and foreign alphabets.
[SC2020] `tr` replaces sets of chars, not words (mentioned due to duplicates).
[SC2021] Don't use `[]` around ranges in `tr`, it replaces literal square brackets.
[SC2022] Note that unlike globs, `o*` here matches `ooo` but not `oscar`.
[SC2023] The shell may override `time` as seen in man time(1). Use `command time ..` for that one.
[SC2024] `sudo` doesn't affect redirects. Use `..| sudo tee file`
[SC2025] Make sure all escape sequences are enclosed in `\[..\]` to prevent line wrapping issues.
[SC2026] This word is outside of quotes. Did you intend to `'nest '"'single quotes'"'` instead'?
[SC2027] The surrounding quotes actually unquote this. Remove or escape them.
[SC2028] `echo` won't expand escape sequences. Consider `printf`.
[SC2029] Note that, unescaped, this expands on the client side.
[SC2030] Modification of var is local (to subshell caused by pipeline).
[SC2031] var was modified in a subshell. That change might be lost.
[SC2032] This function can't be invoked via su on line 42.
[SC2033] Shell functions can't be passed to external commands. Use separate script or sh -c.
[SC2034] foo appears unused. Verify it or export it.
[SC2035] Use `./*glob*` or `-- *glob*` so names with dashes won't become options.
[SC2036] If you wanted to assign the output of the pipeline, use `a=$(b | c)` .
[SC2037] To assign the output of a command, use `var=$(cmd)` .
[SC2038] Use `-print0`/`-0` or `find -exec +` to allow for non-alphanumeric filenames.
[SC2039] In POSIX sh, *something* is undefined.
[SC2040] `#!/bin/sh` was specified, so ____ is not supported, even when sh is actually bash.
[SC2041] This is a literal string. To run as a command, use `$(..)` instead of `'..'` .
[SC2042] Use spaces, not commas, to separate loop elements.
[SC2043] This loop will only ever run once for a constant value. Did you perhaps mean to loop over `dir/*`, `$var` or `$(cmd)`?
[SC2044] For loops over find output are fragile. Use `find -exec` or a `while read` loop.
[SC2045] Iterating over ls output is fragile. Use globs.
[SC2046] Quote this to prevent word splitting.
[SC2048] Use `"$@"` (with quotes) to prevent whitespace problems.
[SC2049] `=~` is for regex, but this looks like a glob. Use `=` instead.
[SC2050] This expression is constant. Did you forget the `$` on a variable?
[SC2051] Bash doesn't support variables in brace range expansions.
[SC2053] Quote the rhs of `=` in `[[ ]]` to prevent glob matching.
[SC2054] Use spaces, not commas, to separate array elements.
[SC2055] You probably wanted `&&` here, otherwise it's always true.
[SC2056] You probably wanted `&&` here
[SC2057] Unknown binary operator.
[SC2058] Unknown unary operator.
[SC2059] Don't use variables in the `printf` format string. Use `printf "..%s.." "$foo"`.
[SC2060] Quote parameters to tr to prevent glob expansion.
[SC2061] Quote the parameter to `-name` so the shell won't interpret it.
[SC2062] Quote the grep pattern so the shell won't interpret it.
[SC2063] Grep uses regex, but this looks like a glob.
[SC2064] Use single quotes, otherwise this expands now rather than when signalled.
[SC2065] This is interpreted as a shell file redirection, not a comparison.
[SC2066] Since you double quoted this, it will not word split, and the loop will only run once.
[SC2067] Missing `;` or `+` terminating `-exec`. You can't use `|`/`||`/`&&`, and `;` has to be a separate, quoted argument.
[SC2068] Double quote array expansions to avoid re-splitting elements.
[SC2069] To redirect stdout+stderr, `2>&1` must be last (or use `{ cmd > file; } 2>&1` to clarify).
[SC2070] `-n` doesn't work with unquoted arguments. Quote or use `[[ ]]`.
[SC2071] `>` is for string comparisons. Use `-gt` instead.
[SC2072] Decimals are not supported. Either use integers only, or use `bc` or `awk` to compare.
[SC2073] Escape `\<` to prevent it redirecting (or switch to `[[ .. ]]`).
[SC2074] Can't use `=~` in `[ ]`. Use `[[..]]` instead.
[SC2075] Escaping `\<` is required in `[..]`, but invalid in `[[..]]`
[SC2076] Don't quote rhs of `=~`, it'll match literally rather than as a regex.
[SC2077] You need spaces around the comparison operator.
[SC2078] This expression is constant. Did you forget a `$` somewhere?
[SC2079] `(( ))` doesn't support decimals. Use `bc` or `awk`.
[SC2080] Numbers with leading 0 are considered octal.
[SC2081] `[ .. ]` can't match globs. Use `[[ .. ]]` or grep.
[SC2082] To expand via indirection, use `name="foo$n"; echo "${!name}"`.
[SC2083] Don't add spaces after the slash in `./file`.
[SC2084] Remove `$` or use `_=$((expr))` to avoid executing output.
[SC2086] Double quote to prevent globbing and word splitting.
[SC2087] Quote `EOF` to make here document expansions happen on the server side rather than on the client.
[SC2088] Tilde does not expand in quotes. Use `$HOME`.
[SC2089] Quotes/backslashes will be treated literally. Use an array.
[SC2090] Quotes/backslashes in this variable will not be respected.
[SC2091] Remove surrounding `$()` to avoid executing output (or use `eval` if intentional).
[SC2092] Remove backticks to avoid executing output.
[SC2093] Remove `exec ` if script should continue after this command.
[SC2094] Make sure not to read and write the same file in the same pipeline.
[SC2095] Use `ssh -n` to prevent ssh from swallowing stdin.
[SC2096] On most OS, shebangs can only specify a single parameter.
[SC2097] This assignment is only seen by the forked process.
[SC2098] This expansion will not see the mentioned assignment.
[SC2099] Use `$((..))` for arithmetics, e.g. `i=$((i + 2))`
[SC2100] Use `$((..))` for arithmetics, e.g. `i=$((i + 2))`
[SC2101] Named class needs outer `[]`, e.g. `[[:digit:]]`.
[SC2102] Ranges can only match single chars (mentioned due to duplicates).
[SC2103] Use a `( subshell )` to avoid having to `cd` back.
[SC2104] In functions, use `return` instead of `break`.
[SC2105] `break` is only valid in loops
[SC2106] This only exits the subshell caused by the pipeline.
[SC2107] Instead of `[ a && b ]`, use `[ a ] && [ b ]`.
[SC2108] In `[[..]]`, use `&&` instead of `-a`.
[SC2109] Instead of `[ a || b ]`, use `[ a ] || [ b ]`.
[SC2110] In `[[..]]`, use `||` instead of `-o`.
[SC2111] ksh does not allow `function` keyword and `()` at the same time.
[SC2112] `function` keyword is non-standard. Delete it.
[SC2113] `function` keyword is non-standard. Use `foo()` instead of `function foo`.
[SC2114] Warning: deletes a system directory.
[SC2115] Use `"${var:?}"` to ensure this never expands to `/*` .
[SC2116] Useless echo? Instead of `cmd $(echo foo)`, just use `cmd foo`.
[SC2117] To run commands as another user, use `su -c` or `sudo`.
[SC2118] Ksh does not support `|&`. Use `2>&1 |`
[SC2119] Use `foo "$@"` if function's `$1` should mean script's `$1`.
[SC2120] foo references arguments, but none are ever passed.
[SC2121] To assign a variable, use just `var=value`, not `set ..`.
[SC2122] `>=` is not a valid operator. Use `! a < b` instead.
[SC2123] `PATH` is the shell search path. Use another name.
[SC2124] Assigning an array to a string! Assign as array, or use `*` instead of `@` to concatenate.
[SC2125] Brace expansions and globs are literal in assignments. Quote it or use an array.
[SC2126] Consider using `grep -c` instead of `grep | wc`
[SC2127] To use `${ ..; }`, specify `#!/usr/bin/env ksh`.
[SC2128] Expanding an array without an index only gives the element in the index 0.
[SC2129] Consider using `{ cmd1; cmd2; } >> file` instead of individual redirects.
[SC2130] `-eq` is for integer comparisons. Use `=` instead.
[SC2139] This expands when defined, not when used. Consider escaping.
[SC2140] Word is of the form `"A"B"C"` (B indicated). Did you mean `"ABC"` or `"A\\"B\\"C"`?
[SC2141] Did you mean `IFS=$'\t'` ?
[SC2142] Aliases can't use positional parameters. Use a function.
[SC2143] Use `grep -q` instead of comparing output with `[ -n .. ]`.
[SC2144] `-e` doesn't work with globs. Use a `for` loop.
[SC2145] Argument mixes string and array. Use `*` or separate argument.
[SC2146] This action ignores everything before the `-o`. Use `\( \)` to group.
[SC2147] Literal tilde in PATH works poorly across programs.
[SC2148] Tips depend on target shell and yours is unknown. Add a shebang.
[SC2149] Remove `$`/`${}` for numeric index, or escape it for string.
[SC2150] `-exec` does not automatically invoke a shell. Use `-exec sh -c ..` for that.
[SC2151] Only one integer 0-255 can be returned. Use stdout for other data.
[SC2152] Can only return 0-255. Other data should be written to stdout.
[SC2153] Possible Misspelling: MYVARIABLE may not be assigned. Did you mean MY_VARIABLE?
[SC2154] var is referenced but not assigned.
[SC2155] Declare and assign separately to avoid masking return values.
[SC2156] Injecting filenames is fragile and insecure. Use parameters.
[SC2157] Argument to implicit `-n` is always true due to literal strings.
[SC2158] `[ false ]` is true. Remove the brackets
[SC2159] `[ 0 ]` is true. Use `false` instead.
[SC2160] Instead of `[ true ]`, just use `true`.
[SC2161] Instead of `[ 1 ]`, use `true`.
[SC2162] `read` without `-r` will mangle backslashes.
[SC2163] This does not export `FOO`. Remove `$`/`${}` for that, or use `${var?}` to quiet.
[SC2164] Use `cd ... || exit` in case `cd` fails.
[SC2165] This nested loop overrides the index variable of its parent.
[SC2166] Prefer `[ p ] && [ q ]` as `[ p -a q ]` is not well defined.
[SC2167] This parent loop has its index variable overridden.
[SC2168] `local` is only valid in functions.
[SC2169] In dash, *something* is not supported.
[SC2170] Invalid number for `-eq`. Use `=` to compare as string (or use `$var` to expand as a variable).
[SC2171] Found trailing `]` outside test. Add missing `[` or quote if intentional.
[SC2172] Trapping signals by number is not well defined. Prefer signal names.
[SC2173] SIGKILL/SIGSTOP can not be trapped.
[SC2174] When used with `-p`, `-m` only applies to the deepest directory.
[SC2175] Quote this invalid brace expansion since it should be passed literally to eval
[SC2176] `time` is undefined for pipelines. time single stage or `bash -c` instead.
[SC2177] `time` is undefined for compound commands, use `time sh -c` instead.
[SC2178] Variable was used as an array but is now assigned a string.
[SC2179] Use `array+=("item")` to append items to an array.
[SC2180] Bash does not support multidimensional arrays. Use 1D or associative arrays.
[SC2181] Check exit code directly with e.g. `if mycmd;`, not indirectly with `$?`.
[SC2182] This printf format string has no variables. Other arguments are ignored.
[SC2183] This format string has 2 variables, but is passed 1 arguments.
[SC2184] Quote arguments to unset so they're not glob expanded.
[SC2185] Some finds don't have a default path. Specify `.` explicitly.
[SC2186] tempfile is deprecated. Use mktemp instead.
[SC2187] Ash scripts will be checked as Dash. Add `# shellcheck shell=dash` to silence.
[SC2188] This redirection doesn't have a command. Move to its command (or use `true` as no-op).
[SC2189] You can't have `|` between this redirection and the command it should apply to.
[SC2190] Elements in associative arrays need index, e.g. `array=( [index]=value )` .
[SC2191] The `=` here is literal. To assign by index, use `( [index]=value )` with no spaces. To keep as literal, quote it.
[SC2192] This array element has no value. Remove spaces after `=` or use `""` for empty string.
[SC2193] The arguments to this comparison can never be equal. Make sure your syntax is correct.
[SC2194] This word is constant. Did you forget the `$` on a variable?
[SC2195] This pattern will never match the case statement's word. Double check them.
[SC2196] `egrep` is non-standard and deprecated. Use `grep -E` instead.
[SC2197] `fgrep` is non-standard and deprecated. Use `grep -F` instead.
[SC2198] Arrays don't work as operands in `[ ]`. Use a loop (or concatenate with `*` instead of `@`).
[SC2199] Arrays implicitly concatenate in `[[ ]]`. Use a loop (or explicit `*` instead of `@`).
[SC2200] Brace expansions don't work as operands in `[ ]`. Use a loop.
[SC2201] Brace expansion doesn't happen in `[[ ]]`. Use a loop.
[SC2202] Globs don't work as operands in `[ ]`. Use a loop.
[SC2203] Globs are ignored in `[[ ]]` except right of `=`/`!=`. Use a loop.
[SC2204] `(..)` is a subshell. Did you mean `[ .. ]`, a test expression?
[SC2205] `(..)` is a subshell. Did you mean `[ .. ]`, a test expression?
[SC2206] Quote to prevent word splitting/globbing, or split robustly with mapfile or `read -a`.
[SC2207] Prefer `mapfile` or `read -a` to split command output (or quote to avoid splitting).
[SC2208] Use `[[ ]]` or quote arguments to `-v` to avoid glob expansion.
[SC2209] Use `var=$(command)` to assign output (or quote to assign string).
[SC2210] This is a file redirection. Was it supposed to be a comparison or fd operation?
[SC2211] This is a glob used as a command name. Was it supposed to be in `${..}`, array, or is it missing quoting?
[SC2212] Use `false` instead of empty `[`/`[[` conditionals.
[SC2213] getopts specified `-n`, but it's not handled by this `case`.
[SC2214] This case is not specified by getopts.
[SC2215] This flag is used as a command name. Bad line break or missing `[ .. ]`?
[SC2216] Piping to `rm`, a command that doesn't read stdin. Wrong command or missing `xargs`?
[SC2217] Redirecting to `echo`, a command that doesn't read stdin. Bad quoting or missing `xargs`?
[SC2218] This function is only defined later. Move the definition up.
[SC2219] Instead of `let expr`, prefer `(( expr ))` .
[SC2220] Invalid flags are not handled. Add a `*)` case.
[SC2221] This pattern always overrides a later one.
[SC2222] This pattern never matches because of a previous pattern.
[SC2223] This default assignment may cause DoS due to globbing. Quote it.
[SC2224] This `mv` has no destination. Check the arguments.
[SC2225] This `cp` has no destination. Check the arguments.
[SC2226] This `ln` has no destination. Check the arguments, or specify `.` explicitly.
[SC2227] Redirection applies to the find command itself. Rewrite to work per action (or move to end).
[SC2229] This does not read `foo`. Remove `$`/`${}` for that, or use `${var?}` to quiet.
[SC2230] `which` is non-standard. Use builtin `command -v` instead.
[SC2231] Quote expansions in this `for` loop glob to prevent wordsplitting, e.g. `"$dir"/*.txt` .
[SC2232] Can't use `sudo` with builtins like `cd`. Did you want `sudo sh -c ..` instead?
[SC2233] Remove superfluous `(..)` around condition to avoid subshell overhead.
[SC2234] Remove superfluous `(..)` around test command to avoid subshell overhead.
[SC2235] Use `{ ..; }` instead of `(..)` to avoid subshell overhead.
[SC2236] Use `-n` instead of `! -z`.
[SC2237] Use `[ -n .. ]` instead of `! [ -z .. ]`.
[SC2238] Redirecting to/from command name instead of file. Did you want pipes/xargs (or quote to ignore)?
[SC2239] Ensure the shebang uses the absolute path to the interpreter.
[SC2240] The dot command does not support arguments in sh/dash. Set them as variables.
[SC2241] The exit status can only be one integer 0-255. Use stdout for other data.
[SC2242] Can only exit with status 0-255. Other data should be written to stdout/stderr.
[SC2243] Prefer explicit `-n` to check for output (or run command without `[`/`[[` to check for success)
[SC2244] Prefer explicit `-n` to check non-empty string (or use `=`/`-ne` to check boolean/integer).
[SC2245] -d only applies to the first expansion of this glob. Use a loop to check any/all.
[SC2246] This shebang specifies a directory. Ensure the interpreter is a file.
[SC2247] Flip leading `$` and `"` if this should be a quoted substitution.
[SC2248] Prefer double quoting even when variables don't contain special characters.
[SC2249] Consider adding a default `*)` case, even if it just exits with error.
[SC2250] Prefer putting braces around variable references even when not strictly required.
[SC2251] This `!` is not on a condition and skips errexit. Use `&& exit 1` instead, or make sure `$?` is checked.
[SC2252] You probably wanted `&&` here, otherwise it's always true.
[SC2253] Use `-R` to recurse, or explicitly `a-r` to remove read permissions.
[SC2254] Quote expansions in case patterns to match literally rather than as a glob.
[SC2255] `[ ]` does not apply arithmetic evaluation. Evaluate with `$((..))` for numbers, or use string comparator for strings.
[SC2256] This translated string is the name of a variable. Flip leading `$` and `"` if this should be a quoted substitution.
[SC2257] Arithmetic modifications in command redirections may be discarded. Do them separately.
[SC2259] This redirection overrides piped input. To use both, merge or pass filenames.
[SC2260] This redirection overrides the output pipe. Use `tee` to output to both.
[SC2261] Multiple redirections compete for stdout. Use `cat`, `tee`, or pass filenames instead.
[SC2262] This alias can't be defined and used in the same parsing unit. Use a function instead.
[SC2263] Since they're in the same parsing unit, this command will not refer to the previously mentioned alias.
[SC2264] This function unconditionally re-invokes itself. Missing `command`?
[SC2265] Use && for logical AND. Single & will background and return true.
[SC2266] Use && for logical AND. Single & will background and return true.
[SC2267] GNU `xargs -i` is deprecated in favor of `-I{}`
[SC2268] Avoid x-prefix in comparisons as it no longer serves a purpose.
[SC2269] This variable is assigned to itself, so the assignment does nothing.
[SC2270] To assign positional parameters, use `set -- first second ..` (or use `[ ]` to compare).
[SC2271] For indirection, use arrays, `declare "var$n=value"`, or (for sh) read/eval
[SC2272] Command name contains `==`. For comparison, use `[ "$var" = value ]`.
[SC2273] Sequence of `===`s found. Merge conflict or intended as a commented border?
[SC2274] Command name starts with `===`. Intended as a commented border?
[SC2275] Command name starts with `=`. Bad line break?
[SC2276] This is interpreted as a command name containing `=`. Bad assignment or comparison?
[SC2277] Use `BASH_ARGV0` to assign to `$0` in bash (or use `[ ]` to compare).
[SC2278] `$0` can't be assigned in Ksh (but it does reflect the current function).
[SC2279] `$0` can't be assigned in Dash. This becomes a command name.
[SC2280] `$0` can't be assigned this way, and there is no portable alternative.
[SC2281] Don't use `$`/`${}` on the left side of assignments.
[SC2282] Variable names can't start with numbers, so this is interpreted as a command.
[SC2283] Use `[ ]` to compare values, or remove spaces around `=` to assign (or quote `'='` if literal).
[SC2284] Use `[ x = y ]` to compare values (or quote `'=='` if literal).
[SC2285] Remove spaces around `+=` to assign (or quote `'+='` if literal).
[SC2286] This empty string is interpreted as a command name. Double check syntax (or use 'true' as a no-op).
[SC2287] This is interpreted as a command name ending with '/'. Double check syntax.
[SC2288] This is interpreted as a command name ending with apostrophe. Double check syntax.
[SC2289] This is interpreted as a command name containing a linefeed. Double check syntax.
[SC2290] Remove spaces around = to assign.
[SC2291] Quote repeated spaces to avoid them collapsing into one.
[SC2292] Prefer `[[ ]]` over `[ ]` for tests in Bash/Ksh.
[SC2293] When eval'ing @Q-quoted words, use * rather than @ as the index.
[SC2294] eval negates the benefit of arrays. Drop eval to preserve whitespace/symbols (or eval as string).
[SC2295] Expansions inside `${..}` need to be quoted separately, otherwise they will match as a pattern.
[SC2296] Parameter expansions can't start with `{`. Double check syntax.
[SC2297] Double quotes must be outside `${}`: `${"invalid"}` vs `"${valid}"`.
[SC2298] `${$x}` is invalid. For expansion, use ${x}. For indirection, use arrays, ${!x} or (for sh) eval.
[SC2299] Parameter expansions can't be nested. Use temporary variables.
[SC2300] Parameter expansion can't be applied to command substitutions. Use temporary variables.
[SC2301] Parameter expansion starts with unexpected quotes. Double check syntax.
[SC2302] This loops over values. To loop over keys, use `"${!array[@]}"`.
[SC2303] `i` is an array value, not a key. Use directly or loop over keys instead.
[SC2304] `*` must be escaped to multiply: `\*`. Modern `$((x * y))` avoids this issue.
[SC2305] Quote regex argument to expr to avoid it expanding as a glob.
[SC2306] Escape glob characters in arguments to expr to avoid pathname expansion.
[SC2307] 'expr' expects 3+ arguments but sees 1. Make sure each operator/operand is a separate argument, and escape <>&|.
[SC2308] `expr length` has unspecified results. Prefer `${#var}`.
[SC2309] -eq treats this as a variable. Use = to compare as string (or expand explicitly with $var)
[SC2310] This function is invoked in an 'if' condition so set -e will be disabled. Invoke separately if failures should cause the script to exit.
[SC2311] Bash implicitly disabled set -e for this function invocation because it's inside a command substitution. Add set -e; before it or enable inherit_errexit.
[SC2312] Consider invoking this command separately to avoid masking its return value (or use '|| true' to ignore).
[SC2313] Quote array indices to avoid them expanding as globs.
[SC2314] In bats, `!` does not cause a test failure.
[SC2315] In bats, `!` does not cause a test failure. Fold the `!` into the conditional!
[SC2316] This applies local to the variable named readonly, which is probably not what you want. Use a separate command or the appropriate `declare` options instead.
[SC2317] Command appears to be unreachable. Check usage (or ignore if invoked indirectly).
[SC2318] This assignment is used again in this `declare`, but won't have taken effect. Use two `declare`s.
[SC2319] This `$?` refers to a condition, not a command. Assign to a variable to avoid it being overwritten.
[SC2320] This `$?` refers to echo/printf, not a previous command. Assign to variable to avoid it being overwritten.
[SC2321] Array indices are already arithmetic contexts. Prefer removing the `$((` and `))`.
[SC2322] In arithmetic contexts, `((x))` is the same as `(x)`. Prefer only one layer of parentheses.
[SC2323] `a[(x)]` is the same as `a[x]`. Prefer not wrapping in additional parentheses.
[SC2324] var+=1 will append, not increment. Use (( var += 1 )), declare -i var, or quote number to silence.
[SC2325] Multiple ! in front of pipelines are a bash/ksh extension. Use only 0 or 1.
[SC2326] ! is not allowed in the middle of pipelines. Use command group as in `cmd | { ! cmd; }` if necessary.
[SC2327] This command substitution will be empty because the command's output gets redirected away.
[SC2328] This redirection takes output away from the command substitution.
[SC2329] This function is never invoked. Check usage (or ignored if invoked indirectly).
[SC2330] BusyBox `[[ .. ]]` does not support glob matching. Use a case statement.
[SC3001] In POSIX sh, process substitution is undefined.
[SC3002] In POSIX sh, extglob is undefined.
[SC3003] In POSIX sh, `$'..'` is undefined.
[SC3004] In POSIX sh, $".." is undefined
[SC3005] In POSIX sh, arithmetic for loops are undefined.
[SC3006] In POSIX sh, standalone `((..))` is undefined.
[SC3007] In POSIX sh, `$[..]` in place of `$((..))` is undefined.
[SC3008] In POSIX sh, select loops are undefined.
[SC3009] In POSIX `sh`, brace expansion is undefined.
[SC3010] In POSIX sh, `[[ ]]` is undefined.
[SC3011] In POSIX sh, here-strings are undefined.
[SC3012] In POSIX sh, lexicographical `\<` is undefined.
[SC3013] In POSIX sh, `-nt` is undefined.
[SC3014] In POSIX sh, `==` in place of `=` is undefined.
[SC3015] In POSIX sh, `=~` regex matching is undefined.
[SC3016] In POSIX sh, unary `-v` (in place of `[ -n "${var+x}" ]`) is undefined.
[SC3017] In POSIX sh, unary `-a` in place of `-e` is undefined.
[SC3018] In POSIX sh, `++` is undefined.
[SC3019] In POSIX sh, exponentials are undefined.
[SC3020] In POSIX sh, `&>` is undefined.
[SC3021] In POSIX sh, `>& filename` (as opposed to `>& fd`) is undefined.
[SC3022] In POSIX sh, named file descriptors is undefined.
[SC3023] In POSIX sh, FDs outside of 0-9 are undefined.
[SC3024] In POSIX sh, `+=` is undefined.
[SC3025] In POSIX sh, `/dev/{tcp,udp}` is undefined.
[SC3026] In POSIX sh, `^` in place of `!` in glob bracket expressions is undefined.
[SC3028] In POSIX sh, VARIABLE is undefined.
[SC3029] In POSIX sh, `|&` in place of `2>&1 |` is undefined.
[SC3030] In POSIX sh, arrays are undefined.
[SC3031] In POSIX sh, redirecting from/to globs is undefined.
[SC3032] In POSIX sh, coproc is undefined.
[SC3033] In POSIX sh, naming functions outside [a-zA-Z_][a-zA-Z0-9_]* is undefined.
[SC3034] In POSIX sh, `$(<file)` is undefined.
[SC3035] In POSIX sh, `` `<file` `` is undefined.
[SC3036] In Dash, echo flags besides -n are not supported.
[SC3037] In POSIX sh, echo flags are undefined.
[SC3038] In POSIX sh, exec flags are undefined.
[SC3039] In POSIX sh, `let` is undefined.
[SC3040] In POSIX sh, set option *[name]* is undefined.
[SC3041] In POSIX sh, set flag `-E` is undefined
[SC3042] In POSIX sh, set flag `--default` is undefined
[SC3043] In POSIX sh, `local` is undefined.
[SC3044] In POSIX sh, `declare` is undefined.
[SC3045] In POSIX sh, some-command-with-flag is undefined.
[SC3046] In POSIX sh, `source` in place of `.` is undefined.
[SC3047] In POSIX sh, trapping ERR is undefined.
[SC3048] In POSIX sh, prefixing signal names with 'SIG' is undefined.
[SC3049] In POSIX sh, using lower/mixed case for signal names is undefined.
[SC3050] In POSIX sh, `printf %q` is undefined.
[SC3051] In POSIX sh, `source` in place of `.` is undefined
[SC3052] In POSIX sh, arithmetic base conversion is undefined
[SC3053] In POSIX sh, indirect expansion is undefined.
[SC3054] In POSIX sh, array references are undefined.
[SC3055] In POSIX sh, array key expansion is undefined.
[SC3056] In POSIX sh, name matching prefixes are undefined.
[SC3057] In POSIX sh, string indexing is undefined.
[SC3059] Case modification is not supported in dash and undefined in POSIX sh.
[SC3060] In POSIX sh, string replacement is undefined.
