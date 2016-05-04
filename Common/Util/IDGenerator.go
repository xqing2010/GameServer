package Util
//IDGenerator generator uniq id
type IDGenerator struct{
    id int;
}

//GetUniqID  return the uniq id in runtime
func (idGen *IDGenerator)GetUniqID() int {
    var id = idGen.id;
    idGen.id++;
    return id;
}