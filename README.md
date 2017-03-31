Easy tuning of CP/M for actual RAM size

Author - Alexey V. Voronin @ FoxyLab © 2017

https://www.acdc.foxylab.com

getcpm licensed under the GPL v3.0.

Build:

go build getcpm.go

Use:

getcpm XX YYYYYYYYYYYY    :

XX - actual RAM size (KBytes);
YYYYYYYYYYYY - CP/M serial number (hex) 

Example:

getcpm 32 010203040506

Result:

CPMXXK.SYS file (CCP & BDOS)
