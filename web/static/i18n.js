// Internationalization system for CV Generator
const i18n = {
    currentLang: 'es',

    translations: {
        es: {
            app: {
                title: 'CV Generator'
            },
            actions: {
                preview: 'Vista Previa',
                export: 'Exportar PDF',
                add: 'Agregar',
                remove: 'Eliminar'
            },
            sections: {
                personal: 'Información Personal',
                experience: 'Experiencia Laboral',
                education: 'Educación',
                skills: 'Habilidades',
                languages: 'Idiomas'
            },
            fields: {
                fullName: 'Nombre Completo *',
                email: 'Email *',
                phone: 'Teléfono',
                location: 'Ubicación',
                linkedin: 'LinkedIn',
                github: 'GitHub',
                website: 'Sitio Web',
                summary: 'Resumen Profesional',
                company: 'Empresa *',
                position: 'Cargo *',
                startDate: 'Fecha de Inicio',
                endDate: 'Fecha de Fin',
                description: 'Descripción',
                institution: 'Institución *',
                degree: 'Título/Grado *',
                skillName: 'Habilidad',
                skillLevel: 'Nivel',
                language: 'Idioma'
            },
            placeholders: {
                fullName: 'Tu nombre completo',
                email: 'tu@email.com',
                phone: '+1 234 567 8900',
                location: 'Ciudad, País',
                linkedin: 'https://linkedin.com/in/tu-perfil',
                github: 'https://github.com/tu-usuario',
                website: 'https://tu-sitio.com',
                summary: 'Describe brevemente tu experiencia y objetivos profesionales...',
                company: 'Nombre de la empresa',
                position: 'Tu cargo o posición',
                description: 'Describe tus responsabilidades y logros...',
                institution: 'Universidad, instituto, etc.',
                degree: 'Carrera, certificación, etc.',
                skillName: 'Habilidad (ej: JavaScript, Liderazgo)',
                language: 'Idioma (ej: Español - Nativo, Inglés - Avanzado)'
            },
            preview: {
                title: 'Vista Previa del CV'
            },
            experience: {
                title: 'Experiencia',
                number: 'Experiencia'
            },
            education: {
                title: 'Educación',
                number: 'Educación'
            },
            skills: {
                levels: {
                    '': 'Nivel',
                    'Básico': 'Básico',
                    'Intermedio': 'Intermedio',
                    'Avanzado': 'Avanzado',
                    'Experto': 'Experto'
                }
            },
            notes: {
                currentJob: 'Deja vacío si es tu trabajo actual',
                currentStudy: 'Deja vacío si aún estudias',
                achievements: 'Menciones honoríficas, proyectos relevantes...'
            }
        },
        en: {
            app: {
                title: 'CV Generator'
            },
            actions: {
                preview: 'Preview',
                export: 'Export PDF',
                add: 'Add',
                remove: 'Remove'
            },
            sections: {
                personal: 'Personal Information',
                experience: 'Work Experience',
                education: 'Education',
                skills: 'Skills',
                languages: 'Languages'
            },
            fields: {
                fullName: 'Full Name *',
                email: 'Email *',
                phone: 'Phone',
                location: 'Location',
                linkedin: 'LinkedIn',
                github: 'GitHub',
                website: 'Website',
                summary: 'Professional Summary',
                company: 'Company *',
                position: 'Position *',
                startDate: 'Start Date',
                endDate: 'End Date',
                description: 'Description',
                institution: 'Institution *',
                degree: 'Degree *',
                skillName: 'Skill',
                skillLevel: 'Level',
                language: 'Language'
            },
            placeholders: {
                fullName: 'Your full name',
                email: 'your@email.com',
                phone: '+1 234 567 8900',
                location: 'City, Country',
                linkedin: 'https://linkedin.com/in/your-profile',
                github: 'https://github.com/your-username',
                website: 'https://your-website.com',
                summary: 'Briefly describe your experience and professional goals...',
                company: 'Company name',
                position: 'Your role or position',
                description: 'Describe your responsibilities and achievements...',
                institution: 'University, institute, etc.',
                degree: 'Degree, certification, etc.',
                skillName: 'Skill (e.g: JavaScript, Leadership)',
                language: 'Language (e.g: Spanish - Native, English - Advanced)'
            },
            preview: {
                title: 'CV Preview'
            },
            experience: {
                title: 'Experience',
                number: 'Experience'
            },
            education: {
                title: 'Education',
                number: 'Education'
            },
            skills: {
                levels: {
                    '': 'Level',
                    'Básico': 'Basic',
                    'Intermedio': 'Intermediate',
                    'Avanzado': 'Advanced',
                    'Experto': 'Expert'
                }
            },
            notes: {
                currentJob: 'Leave empty if this is your current job',
                currentStudy: 'Leave empty if you are still studying',
                achievements: 'Honors, relevant projects...'
            }
        }
    },

    // Initialize i18n system
    init() {
        // Get saved language or default to Spanish
        this.currentLang = localStorage.getItem('cv-generator-lang') || 'es';
        document.documentElement.setAttribute('data-lang', this.currentLang);
        document.documentElement.setAttribute('lang', this.currentLang);

        // Set language selector
        const langSelect = document.getElementById('language-select');
        if (langSelect) {
            langSelect.value = this.currentLang;
        }

        this.updateTexts();
    },

    // Change language
    changeLanguage(lang) {
        this.currentLang = lang;
        localStorage.setItem('cv-generator-lang', lang);
        document.documentElement.setAttribute('data-lang', lang);
        document.documentElement.setAttribute('lang', lang);
        this.updateTexts();

        // Update page title
        document.title = this.t('app.title') + ' - ' + (lang === 'es'
            ? 'Generador de CV Profesional Minimalista'
            : 'Professional Minimalist CV Generator');
    },

    // Get translation
    t(key) {
        const keys = key.split('.');
        let value = this.translations[this.currentLang];

        for (const k of keys) {
            if (value && value[k]) {
                value = value[k];
            } else {
                return key; // Return key if translation not found
            }
        }

        return value;
    },

    // Update all texts in the page
    updateTexts() {
        // Update elements with data-i18n attribute
        document.querySelectorAll('[data-i18n]').forEach(element => {
            const key = element.getAttribute('data-i18n');
            element.textContent = this.t(key);
        });

        // Update placeholders
        document.querySelectorAll('[data-i18n-placeholder]').forEach(element => {
            const key = element.getAttribute('data-i18n-placeholder');
            element.setAttribute('placeholder', this.t(key));
        });

        // Update meta tags
        const description = document.querySelector('meta[name="description"]');
        if (description) {
            const newDesc = this.currentLang === 'es'
                ? 'Generador de CV profesional minimalista. Crea currículums elegantes en formato PDF con diseño monocromático estilo Notion. Gratis y fácil de usar.'
                : 'Professional minimalist CV generator. Create elegant resumes in PDF format with monochromatic Notion-style design. Free and easy to use.';
            description.setAttribute('content', newDesc);
        }

        // Update existing dynamic content
        this.updateDynamicContent();
    },

    // Update dynamic content (experience, education, etc.)
    updateDynamicContent() {
        // Update experience items
        document.querySelectorAll('#experience-container .item-number').forEach((element, index) => {
            element.textContent = `${this.t('experience.number')} ${index + 1}`;
        });

        // Update education items
        document.querySelectorAll('#education-container .item-number').forEach((element, index) => {
            element.textContent = `${this.t('education.number')} ${index + 1}`;
        });

        // Update skill level options
        document.querySelectorAll('.skill-level option').forEach(option => {
            const value = option.value;
            if (this.t('skills.levels')[value] !== undefined) {
                option.textContent = this.t('skills.levels')[value];
            }
        });

        // Update notes/hints
        document.querySelectorAll('small').forEach(small => {
            const text = small.textContent.trim();
            if (text.includes('trabajo actual') || text.includes('current job')) {
                small.textContent = this.t('notes.currentJob');
            } else if (text.includes('aún estudias') || text.includes('still studying')) {
                small.textContent = this.t('notes.currentStudy');
            } else if (text.includes('proyectos relevantes') || text.includes('relevant projects')) {
                small.textContent = this.t('notes.achievements');
            }
        });
    }
};

// Global function for language change
function changeLanguage(lang) {
    i18n.changeLanguage(lang);
}

// Initialize when DOM is loaded
document.addEventListener('DOMContentLoaded', function () {
    i18n.init();
});
