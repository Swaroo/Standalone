bool isValid(char * s){
    int len = strlen(s);
    if (len <= 0){
        return true;
    }
    char check[len];
    int pos = 0;
    for (int i =0 ; i<len;i++) {
        switch (s[i]){
            case '(':
            case '{':
            case '[':check[pos++] = s[i];
                break;
            case ')':
                if ((pos <=0) || (check[pos-1]!='(')){
                return false;
            }
            pos--;
                break;
            case '}':if ((pos <=0) || (check[pos-1]!='{')){
                return false;
            }
            pos--;
                break;
            case ']':if ((pos <=0) || (check[pos-1]!='[')){
                return false;
            }
            pos--;
                break;
        }
    } 
    if (pos != 0){
        return false;
    }
    return true;
}
