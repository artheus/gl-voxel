#version 330 core
out vec4 FragColor;

uniform sampler2D tex;

in vec2 fragTexCoord;

in vec3 Normal;
in vec3 FragPos;

void main()
{
    vec3 lightPos = vec3(0,80,0);
    vec3 viewPos = vec3(50,50,50);
    vec3 lightColor = vec3(1.0, 1.0, 1.0);
    vec4 objectColor = texture(tex, vec2(fragTexCoord.x, 1-fragTexCoord.y));
    if (objectColor == vec4(0,0,0,0)) {
        discard;
    }

    if (objectColor.w > 0 && objectColor.w < 1) {
        FragColor = objectColor;
        return;
        //objectColor.w = 1;
    }

    // ambient
    float ambientStrength = 0.5;
    vec3 ambient = ambientStrength * lightColor;

    // diffuse
    vec3 norm = normalize(Normal);
    vec3 lightDir = normalize(lightPos - FragPos);
    float diff = max(dot(norm, lightDir), 0.0);
    vec3 diffuse = diff * lightColor;

    // specular
    float specularStrength = 0.0;
    vec3 viewDir = normalize(viewPos - FragPos);
    vec3 reflectDir = reflect(-lightDir, norm);
    float spec = pow(max(dot(viewDir, reflectDir), 0.0), 32);
    vec3 specular = specularStrength * spec * lightColor;

    vec4 lightResult = vec4(ambient + diffuse + specular, 1.0);
    vec4 result = lightResult * objectColor;
    FragColor = result;
}