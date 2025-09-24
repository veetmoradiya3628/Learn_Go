// Excersize for characters 
class SqueakyClean {
    static String clean(String identifier) {
        StringBuilder cleanedIdentifier = new StringBuilder();
        boolean nextIsUpperCase = false;
        
        for (char ch : identifier.toCharArray()) {
            if (ch == '-') {
                nextIsUpperCase = true;
                continue;
            }

            if (Character.isWhitespace(ch)) {
                cleanedIdentifier.append('_');
                nextIsUpperCase = false;
                continue;
            }

            if (Character.isLetter(ch)) {
                if (nextIsUpperCase) {
                    cleanedIdentifier.append(Character.toUpperCase(ch));
                    nextIsUpperCase = false;
                } else {
                    cleanedIdentifier.append(ch);
                }
                continue;
            }

            switch (ch) {
                case '4':
                    cleanedIdentifier.append('a');
                    break;
                case '3':
                    cleanedIdentifier.append('e');
                    break;
                case '0':
                    cleanedIdentifier.append('o');
                    break;
                case '1':
                    cleanedIdentifier.append('l');
                    break;
                case '7':
                    cleanedIdentifier.append('t');
                    break;
                default:
                    break;
            }
        }
        return cleanedIdentifier.toString();
    }
}
