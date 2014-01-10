portsetup
=========

Provides and manage free local ports to be used in connecting to a specific remote port over proxy. This is custom made to be used with  TunnelSetup, TAPmanager and tapdaemon

**USAGE:**

Change portsetup.cfg to correct data. If not user is not set, it will use UNIX account name.
*The port used will probably be dynamic. Add -p before -e parameter.*

**Allocate TAP / IP:**

portsetup -e allocate    
*(To see all data allocated: portsetup -v -e allocate   )*

**Get IP address:**

portsetup -e ip

*Example:*
<pre>
set IP=`portsetup -p $MANAGER_PORT -e ip`
echo $IP
</pre>

**Get Port:**

portsetup -e port

*Example:*
<pre>
set PORT=`portsetup -p $MANAGER_PORT -e port`
echo $PORT
</pre>

**Remove TAP allocation**

portsetup -p $MANAGER_PORT -e remove

