// Global state for dynamic data
let experienceData = [];
let educationData = [];
let skillsData = [];
let languagesData = [];

// Initialize the application
document.addEventListener('DOMContentLoaded', function () {
    // Add initial items
    addExperience();
    addEducation();
    addSkill();
    addLanguage();

    // Setup form submission
    document.getElementById('cv-form').addEventListener('submit', function (e) {
        updateHiddenInputs();
    });
});

// Experience Management
function addExperience() {
    const container = document.getElementById('experience-container');
    const index = experienceData.length;

    experienceData.push({
        company: '',
        position: '',
        startDate: '',
        endDate: '',
        description: ''
    });

    const itemHtml = `
        <div class="dynamic-item" data-index="${index}">
            <div class="item-header">
                <span class="item-number">Experiencia ${index + 1}</span>
                <div class="item-actions">
                    <button type="button" class="btn btn-danger" onclick="removeExperience(${index})">
                        <i class="fas fa-trash"></i>
                        Eliminar
                    </button>
                </div>
            </div>
            <div class="form-grid">
                <div class="form-group">
                    <label>Empresa *</label>
                    <input type="text" required placeholder="Nombre de la empresa" 
                           onchange="updateExperience(${index}, 'company', this.value)">
                </div>
                <div class="form-group">
                    <label>Cargo *</label>
                    <input type="text" required placeholder="Tu cargo o posición" 
                           onchange="updateExperience(${index}, 'position', this.value)">
                </div>
                <div class="form-group">
                    <label>Fecha de Inicio</label>
                    <input type="month" onchange="updateExperience(${index}, 'startDate', this.value)">
                </div>
                <div class="form-group">
                    <label>Fecha de Fin</label>
                    <input type="month" onchange="updateExperience(${index}, 'endDate', this.value)">
                    <small style="color: var(--text-tertiary); font-size: 12px;">Deja vacío si es tu trabajo actual</small>
                </div>
                <div class="form-group full-width">
                    <label>Descripción</label>
                    <textarea rows="3" placeholder="Describe tus responsabilidades y logros..." 
                              onchange="updateExperience(${index}, 'description', this.value)"></textarea>
                </div>
            </div>
        </div>
    `;

    container.insertAdjacentHTML('beforeend', itemHtml);
}

function updateExperience(index, field, value) {
    if (experienceData[index]) {
        experienceData[index][field] = value;
    }
}

function removeExperience(index) {
    const container = document.getElementById('experience-container');
    const item = container.querySelector(`[data-index="${index}"]`);
    if (item) {
        item.remove();
        experienceData.splice(index, 1);
        updateExperienceIndices();
    }
}

function updateExperienceIndices() {
    const container = document.getElementById('experience-container');
    const items = container.querySelectorAll('.dynamic-item');
    items.forEach((item, newIndex) => {
        item.setAttribute('data-index', newIndex);
        item.querySelector('.item-number').textContent = `Experiencia ${newIndex + 1}`;

        // Update onclick handlers
        const inputs = item.querySelectorAll('input, textarea');
        inputs.forEach(input => {
            const onchange = input.getAttribute('onchange');
            if (onchange) {
                input.setAttribute('onchange', onchange.replace(/updateExperience\(\d+,/, `updateExperience(${newIndex},`));
            }
        });

        const removeBtn = item.querySelector('.btn-danger');
        if (removeBtn) {
            removeBtn.setAttribute('onclick', `removeExperience(${newIndex})`);
        }
    });
}

// Education Management
function addEducation() {
    const container = document.getElementById('education-container');
    const index = educationData.length;

    educationData.push({
        institution: '',
        degree: '',
        startDate: '',
        endDate: '',
        description: ''
    });

    const itemHtml = `
        <div class="dynamic-item" data-index="${index}">
            <div class="item-header">
                <span class="item-number">Educación ${index + 1}</span>
                <div class="item-actions">
                    <button type="button" class="btn btn-danger" onclick="removeEducation(${index})">
                        <i class="fas fa-trash"></i>
                        Eliminar
                    </button>
                </div>
            </div>
            <div class="form-grid">
                <div class="form-group">
                    <label>Institución *</label>
                    <input type="text" required placeholder="Universidad, instituto, etc." 
                           onchange="updateEducation(${index}, 'institution', this.value)">
                </div>
                <div class="form-group">
                    <label>Título/Grado *</label>
                    <input type="text" required placeholder="Carrera, certificación, etc." 
                           onchange="updateEducation(${index}, 'degree', this.value)">
                </div>
                <div class="form-group">
                    <label>Fecha de Inicio</label>
                    <input type="month" onchange="updateEducation(${index}, 'startDate', this.value)">
                </div>
                <div class="form-group">
                    <label>Fecha de Fin</label>
                    <input type="month" onchange="updateEducation(${index}, 'endDate', this.value)">
                    <small style="color: var(--text-tertiary); font-size: 12px;">Deja vacío si aún estudias</small>
                </div>
                <div class="form-group full-width">
                    <label>Descripción</label>
                    <textarea rows="2" placeholder="Menciones honoríficas, proyectos relevantes..." 
                              onchange="updateEducation(${index}, 'description', this.value)"></textarea>
                </div>
            </div>
        </div>
    `;

    container.insertAdjacentHTML('beforeend', itemHtml);
}

function updateEducation(index, field, value) {
    if (educationData[index]) {
        educationData[index][field] = value;
    }
}

function removeEducation(index) {
    const container = document.getElementById('education-container');
    const item = container.querySelector(`[data-index="${index}"]`);
    if (item) {
        item.remove();
        educationData.splice(index, 1);
        updateEducationIndices();
    }
}

function updateEducationIndices() {
    const container = document.getElementById('education-container');
    const items = container.querySelectorAll('.dynamic-item');
    items.forEach((item, newIndex) => {
        item.setAttribute('data-index', newIndex);
        item.querySelector('.item-number').textContent = `Educación ${newIndex + 1}`;

        const inputs = item.querySelectorAll('input, textarea');
        inputs.forEach(input => {
            const onchange = input.getAttribute('onchange');
            if (onchange) {
                input.setAttribute('onchange', onchange.replace(/updateEducation\(\d+,/, `updateEducation(${newIndex},`));
            }
        });

        const removeBtn = item.querySelector('.btn-danger');
        if (removeBtn) {
            removeBtn.setAttribute('onclick', `removeEducation(${newIndex})`);
        }
    });
}

// Skills Management
function addSkill() {
    const container = document.getElementById('skills-container');
    const index = skillsData.length;

    skillsData.push({
        name: '',
        level: ''
    });

    const itemHtml = `
        <div class="skill-item" data-index="${index}">
            <input type="text" placeholder="Habilidad (ej: JavaScript, Liderazgo)" 
                   onchange="updateSkill(${index}, 'name', this.value)">
            <select class="skill-level" onchange="updateSkill(${index}, 'level', this.value)">
                <option value="">Nivel</option>
                <option value="Básico">Básico</option>
                <option value="Intermedio">Intermedio</option>
                <option value="Avanzado">Avanzado</option>
                <option value="Experto">Experto</option>
            </select>
            <button type="button" class="btn btn-danger" onclick="removeSkill(${index})">
                <i class="fas fa-trash"></i>
            </button>
        </div>
    `;

    container.insertAdjacentHTML('beforeend', itemHtml);
}

function updateSkill(index, field, value) {
    if (skillsData[index]) {
        skillsData[index][field] = value;
    }
}

function removeSkill(index) {
    const container = document.getElementById('skills-container');
    const item = container.querySelector(`[data-index="${index}"]`);
    if (item) {
        item.remove();
        skillsData.splice(index, 1);
        updateSkillIndices();
    }
}

function updateSkillIndices() {
    const container = document.getElementById('skills-container');
    const items = container.querySelectorAll('.skill-item');
    items.forEach((item, newIndex) => {
        item.setAttribute('data-index', newIndex);

        const inputs = item.querySelectorAll('input, select');
        inputs.forEach(input => {
            const onchange = input.getAttribute('onchange');
            if (onchange) {
                input.setAttribute('onchange', onchange.replace(/updateSkill\(\d+,/, `updateSkill(${newIndex},`));
            }
        });

        const removeBtn = item.querySelector('.btn-danger');
        if (removeBtn) {
            removeBtn.setAttribute('onclick', `removeSkill(${newIndex})`);
        }
    });
}

// Languages Management
function addLanguage() {
    const container = document.getElementById('languages-container');
    const index = languagesData.length;

    languagesData.push('');

    const itemHtml = `
        <div class="language-item" data-index="${index}">
            <input type="text" placeholder="Idioma (ej: Español - Nativo, Inglés - Avanzado)" 
                   onchange="updateLanguage(${index}, this.value)">
            <button type="button" class="btn btn-danger" onclick="removeLanguage(${index})">
                <i class="fas fa-trash"></i>
            </button>
        </div>
    `;

    container.insertAdjacentHTML('beforeend', itemHtml);
}

function updateLanguage(index, value) {
    languagesData[index] = value;
}

function removeLanguage(index) {
    const container = document.getElementById('languages-container');
    const item = container.querySelector(`[data-index="${index}"]`);
    if (item) {
        item.remove();
        languagesData.splice(index, 1);
        updateLanguageIndices();
    }
}

function updateLanguageIndices() {
    const container = document.getElementById('languages-container');
    const items = container.querySelectorAll('.language-item');
    items.forEach((item, newIndex) => {
        item.setAttribute('data-index', newIndex);

        const input = item.querySelector('input');
        if (input) {
            input.setAttribute('onchange', `updateLanguage(${newIndex}, this.value)`);
        }

        const removeBtn = item.querySelector('.btn-danger');
        if (removeBtn) {
            removeBtn.setAttribute('onclick', `removeLanguage(${newIndex})`);
        }
    });
}

// Update hidden inputs before form submission
function updateHiddenInputs() {
    // Filter out empty entries
    const filteredExperience = experienceData.filter(exp => exp.company && exp.position);
    const filteredEducation = educationData.filter(edu => edu.institution && edu.degree);
    const filteredSkills = skillsData.filter(skill => skill.name);
    const filteredLanguages = languagesData.filter(lang => lang.trim());

    document.getElementById('experience-data').value = JSON.stringify(filteredExperience);
    document.getElementById('education-data').value = JSON.stringify(filteredEducation);
    document.getElementById('skills-data').value = JSON.stringify(filteredSkills);
    document.getElementById('languages-data').value = JSON.stringify(filteredLanguages);
}

// Preview functionality
function previewCV() {
    updateHiddenInputs();

    const formData = new FormData(document.getElementById('cv-form'));
    const personalInfo = {
        fullName: formData.get('fullName'),
        email: formData.get('email'),
        phone: formData.get('phone'),
        location: formData.get('location'),
        linkedin: formData.get('linkedin'),
        github: formData.get('github'),
        website: formData.get('website'),
        summary: formData.get('summary')
    };

    const experience = JSON.parse(formData.get('experience') || '[]');
    const education = JSON.parse(formData.get('education') || '[]');
    const skills = JSON.parse(formData.get('skills') || '[]');
    const languages = JSON.parse(formData.get('languages') || '[]');

    generatePreview(personalInfo, experience, education, skills, languages);
    showModal();
}

function generatePreview(personalInfo, experience, education, skills, languages) {
    const previewContent = document.getElementById('preview-content');

    let html = `
        <div class="cv-preview">
            <h1>${personalInfo.fullName || 'Tu Nombre'}</h1>
            <div class="contact-info">
    `;

    const contacts = [];
    if (personalInfo.email) contacts.push(personalInfo.email);
    if (personalInfo.phone) contacts.push(personalInfo.phone);
    if (personalInfo.location) contacts.push(personalInfo.location);
    if (personalInfo.linkedin) contacts.push(`LinkedIn: ${personalInfo.linkedin}`);
    if (personalInfo.github) contacts.push(`GitHub: ${personalInfo.github}`);
    if (personalInfo.website) contacts.push(`Website: ${personalInfo.website}`);

    html += contacts.join(' | ');
    html += '</div>';

    if (personalInfo.summary) {
        html += `
            <div class="section">
                <h2>Resumen</h2>
                <p>${personalInfo.summary}</p>
            </div>
        `;
    }

    if (experience.length > 0) {
        html += `
            <div class="section">
                <h2>Experiencia</h2>
        `;
        experience.forEach(exp => {
            html += `
                <div class="experience-item">
                    <div class="item-title">${exp.position} en ${exp.company}</div>
                    <div class="item-subtitle">${exp.startDate} - ${exp.endDate || 'Presente'}</div>
                    ${exp.description ? `<div class="item-description">${exp.description}</div>` : ''}
                </div>
            `;
        });
        html += '</div>';
    }

    if (education.length > 0) {
        html += `
            <div class="section">
                <h2>Educación</h2>
        `;
        education.forEach(edu => {
            html += `
                <div class="education-item">
                    <div class="item-title">${edu.degree} - ${edu.institution}</div>
                    <div class="item-subtitle">${edu.startDate} - ${edu.endDate || 'Presente'}</div>
                    ${edu.description ? `<div class="item-description">${edu.description}</div>` : ''}
                </div>
            `;
        });
        html += '</div>';
    }

    if (skills.length > 0) {
        html += `
            <div class="section">
                <h2>Habilidades</h2>
                <div class="skills-list">
        `;
        const skillTexts = skills.map(skill =>
            skill.level ? `${skill.name} (${skill.level})` : skill.name
        );
        html += skillTexts.join(' • ');
        html += '</div></div>';
    }

    if (languages.length > 0) {
        html += `
            <div class="section">
                <h2>Idiomas</h2>
                <div class="languages-list">${languages.join(' • ')}</div>
            </div>
        `;
    }

    html += '</div>';
    previewContent.innerHTML = html;
}

function showModal() {
    const modal = document.getElementById('preview-modal');
    modal.classList.add('show');
    document.body.style.overflow = 'hidden';
}

function closePreview() {
    const modal = document.getElementById('preview-modal');
    modal.classList.remove('show');
    document.body.style.overflow = '';
}

// Close modal when clicking outside
document.getElementById('preview-modal').addEventListener('click', function (e) {
    if (e.target === this) {
        closePreview();
    }
});

// Close modal with Escape key
document.addEventListener('keydown', function (e) {
    if (e.key === 'Escape') {
        closePreview();
    }
});
