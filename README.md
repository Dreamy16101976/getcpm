<i>getcpm</i> - легкая настройка образа системы CP/M на актуальный размер ОЗУ<br/><br/>
Copyright (C) 2017 Алексей "FoxyLab" Воронин<br/>
Электронная почта:    support@foxylab.com<br/>
Сайт:  https://acdc.foxylab.com<br/>
Это программное обеспечение распространяется под лицензией GPL v3.0.<br/><br/><br/>
Компиляция:<br/>
<b>go build getcpm.go</b><br/>
Использование:<br/>
<b>getcpm XX YYYYYYYYYYYY</b>  :
XX - актуальный размер ОЗУ в Кбайтах;<br/>
YYYYYYYYYYYY - желаемый серийный номер системы CP/M (hex)<br/> 
<br/>
Пример:<br/>
<b>getcpm 32 010203040506</b><br/>
Результат:<br/>
файл CPMXXK.SYS, содержащий CCP и BDOS<br/><br/><br/>
<i>getcpm</i> - easy tuning of CP/M for actual RAM size<br/><br/>
Copyright (C) 2017 Alexey "FoxyLab" Voronin<br/>
Email:    support@foxylab.com<br/>
Website:  https://acdc.foxylab.com<br/>
This software is licensed under the GPL v3.0 License.<br/><br/>
Build:<br/>
<b>go build getcpm.go</b><br/>
Use:<br/>
<b>getcpm XX YYYYYYYYYYYY</b>  :
XX - actual RAM size (KBytes);<br/>
YYYYYYYYYYYY - CP/M serial number (hex)<br/> 
<br/>
Example:<br/>
<b>getcpm 32 010203040506</b><br/>
Result:<br/>
CPMXXK.SYS file (CCP & BDOS)
