char * longestCommonPrefix(char ** strs, int strsSize){
    char * out = "";
    char * str;
    if (strsSize > 0){
    out = strs[0];
    int j = 0;
    for(int i=1; i<strsSize; i++){
        str = strs[i];       
        for(j=0; j<strlen(str);j++){
            if(str[j]!=out[j]){
                out[j]='\0';
                break;
            }
        }
        if (j==0){
            out = "";
            break;
        }else if (j<strlen(out)){
            out[j] = '\0';
        }
    }
    }
    return out;
}
