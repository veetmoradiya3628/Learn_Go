import java.util.ArrayList;
import java.util.List;

public class LanguageList {
    private final List<String> languages = new ArrayList<>();

    public boolean isEmpty() {
        return this.languages.isEmpty();
    }

    public void addLanguage(String language) {
        this.languages.add(language);
    }

    public void removeLanguage(String language) {
        languages.removeIf(item -> item.equals(language));
    }

    public String firstLanguage() {
        return this.languages.get(0);
    }

    public int count() {
        return (int) this.languages.stream().distinct().count();
    }

    public boolean containsLanguage(String language) {
        return this.languages.contains(language);
    }

    public boolean isExciting() {
        return languages.stream()
            .anyMatch(lang -> lang.equalsIgnoreCase("java") || lang.equalsIgnoreCase("kotlin"));
    }
}
