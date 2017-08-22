@If "%1"=="" goto iter
@If "%2"=="" goto %1

@Set args= /I /E /V /C /H /R /K /O /X /B /Y
@xcopy %1	%2		%args%
@goto done

:cleanup
@Echo Do Cleanup
@cd	..\s\
@Del 		Send{{.}}Proxy.dot.go.tmpl
@cd	..\l\
@Del 		Send{{.}}Proxy.dot.go.tmpl
@cd	..\sl\
@Del 		Send{{.}}Proxy.dot.go.tmpl
@cd	..\xs\
@Del 		Send{{.}}Proxy.dot.go.tmpl
@cd	..\xl\
@Del 		Send{{.}}Proxy.dot.go.tmpl

@cd ..\sss\

@goto done

:iter

@Echo xCopy Define.Core.tmpl
@xcopy _Core.nonil		basic\type\Define.Core.tmpl		/Y /Q
@xcopy _Core.tmpl		basic\type\IsUnsafe\Define.Core.tmpl	/Y /Q
@xcopy _Core.merge		basic\type\IsFloat\Define.Core.tmpl	/Y /Q
@xcopy _Core.merge		basic\type\IsInteger\Define.Core.tmpl	/Y /Q
@xcopy _Core.merge		basic\type\IsOrdered\Define.Core.tmpl	/Y /Q
@xcopy _Core.merge		basic\type\IsUnsigned\Define.Core.tmpl	/Y /Q

@Echo Distribute Define.Core.tmpl
@Call %0 *Define.Core.tmpl	..\s\
@Call %0 *Define.Core.tmpl	..\ss\

@Echo Distribute *dot.go.tmpl
@Call %0 *dot.go.tmpl	..\s\
@Call %0 *dot.go.tmpl	..\ss\

@Call %0 *dot.go.tmpl	..\ssss\
@Call %0 *dot.go.tmpl	..\l\
@Call %0 *dot.go.tmpl	..\sl\
@Call %0 *dot.go.tmpl	..\xs\
@Call %0 *dot.go.tmpl	..\xl\

@Echo Call Cleanup
@Call %0 cleanup

@goto done

:done
