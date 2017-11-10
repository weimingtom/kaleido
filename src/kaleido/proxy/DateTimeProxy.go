package proxy

type DateTimeProxy struct { 
	//see https://github.com/weimingtom/mochalua/blob/master/Mochalua/src/com/groundspeak/mochalua/LuaOSLib.java
	//_calendar Calendar
}
		 
func NewDateTimeProxyInit() *DateTimeProxy {
	self := &DateTimeProxy{}
	//this._calendar = Calendar.getInstance(TimeZone.getTimeZone("GMT"));
	return self
}

func NewDateTimeProxyWithYMDHMS(year int, month int, day int, hour int, min int, sec int) *DateTimeProxy {
	self := &DateTimeProxy{}
	//this._calendar = Calendar.getInstance(TimeZone.getTimeZone("GMT"));
	//this._calendar.set(year, month, day, hour, min, sec);
	return self
}

func (self *DateTimeProxy) SetUTCNow() {
	//_calendar = Calendar.getInstance(TimeZone.getTimeZone("UTC"));
}

func (self *DateTimeProxy) SetNow() {
	//_calendar = Calendar.getInstance(TimeZone.getTimeZone("GMT"));  
}

func (self *DateTimeProxy) GetSecond() int {
	//return _calendar.get(Calendar.SECOND);
	return 0
}

func (self *DateTimeProxy) getMinute() int {
	//return _calendar.get(Calendar.MINUTE);
	return 0
}

func (self *DateTimeProxy) GetHour() int {
	//return _calendar.get(Calendar.HOUR_OF_DAY);
	return 0
}

func (self *DateTimeProxy) GetDay() int {
	//return _calendar.get(Calendar.DATE);
	return 0
}

func (self *DateTimeProxy) GetMonth() int {
	//return _calendar.get(Calendar.MONTH) + 1;
	return 0
}

func (self *DateTimeProxy) GetYear() int {
	//return _calendar.get(Calendar.YEAR);
	return 0
}

func (self *DateTimeProxy) GetDayOfWeek() int {
	//return _calendar.get(Calendar.DAY_OF_WEEK);
	return 0
}

func (self *DateTimeProxy) GetDayOfYear() int {
	//return _calendar.get(Calendar.DAY_OF_YEAR);
	return 0
}

//http://www.cnblogs.com/zyw-205520/p/4632490.html
//https://github.com/anonl/luajpp2/blob/master/core/src/main/java/nl/weeaboo/lua2/lib/OsLib.java
func (self *DateTimeProxy) IsDaylightSavingTime() bool {
	//return _calendar.get(Calendar.DST_OFFSET) != 0;
	return false
}

//https://github.com/weimingtom/mochalua/blob/master/Mochalua/src/com/groundspeak/mochalua/LuaOSLib.java
func (self *DateTimeProxy) GetTicks() int64 {
	//return _calendar.getTime().getTime();
	return 0
}

//https://github.com/anonl/luajpp2/blob/master/core/src/main/java/nl/weeaboo/lua2/lib/OsLib.java
//private static final long _t0 = System.currentTimeMillis();
func (self *DateTimeProxy) GetClock() int64 {
	//return (System.currentTimeMillis() - _t0) / 1000.;
	return 0
}