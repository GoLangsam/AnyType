@If "%1"=="" goto iter

@Set args= /I /E /V /C /H /R /K /O /X /B /Y

@xcopy %1	..\s\		%args%
@xcopy %1	..\ss\		%args%

@xcopy %1	..\ssss\	%args%
@xcopy %1	..\l\		%args%
@xcopy %1	..\sl\		%args%
@xcopy %1	..\xs\		%args%
@xcopy %1	..\xl\		%args%

@Del 		..\s\Send{{.}}Proxy.dot.go.tmpl
@Del 		..\l\Send{{.}}Proxy.dot.go.tmpl
@Del 		..\sl\Send{{.}}Proxy.dot.go.tmpl
@Del 		..\xs\Send{{.}}Proxy.dot.go.tmpl
@Del 		..\xl\Send{{.}}Proxy.dot.go.tmpl

@goto done

:iter

xcopy Core.nonil	basic\type\core.tmpl		/Y

xcopy Core.tmpl		basic\type\IsUnsafe\core.tmpl	/Y

xcopy Core.merge	basic\type\IsFloat\core.tmpl	/Y
xcopy Core.merge	basic\type\IsInteger\core.tmpl	/Y
xcopy Core.merge	basic\type\IsOrdered\core.tmpl	/Y
xcopy Core.merge	basic\type\IsUnsigned\core.tmpl	/Y

xcopy Core.basicmerge	basic\type\core.tmpl

@Call %0 *core.tmpl
@Call %0 *dot.go.tmpl

@goto done

:done
